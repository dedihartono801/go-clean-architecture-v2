package transaction

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/repository"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/customstatus"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/dto"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/identifier"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/queue/redis"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"golang.org/x/sync/errgroup"
)

type Service interface {
	Checkout(ctx *fiber.Ctx, input *dto.TransactionCheckoutDto) (*entity.Transaction, int, error)
}

type service struct {
	workerTask              redis.TaskDistributor
	dbTransactionRepository repository.DbTransactionRepository
	transactionRepository   repository.TransactionRepository
	skuRepository           repository.SkuRepository
	validator               validator.Validator
	identifier              identifier.Identifier
	skuMutex                sync.Mutex
}

func NewTransactionService(
	workerTask redis.TaskDistributor,
	dbTransactionRepository repository.DbTransactionRepository,
	transactionRepository repository.TransactionRepository,
	skuRepository repository.SkuRepository,
	validator validator.Validator,
	identifier identifier.Identifier,
) Service {
	return &service{
		workerTask:              workerTask,
		dbTransactionRepository: dbTransactionRepository,
		transactionRepository:   transactionRepository,
		skuRepository:           skuRepository,
		validator:               validator,
		identifier:              identifier,
		skuMutex:                sync.Mutex{},
	}
}

func (s *service) Checkout(ctx *fiber.Ctx, input *dto.TransactionCheckoutDto) (*entity.Transaction, int, error) {

	checkout := dto.TransactionCheckoutDto{
		Items: input.Items,
	}

	if err := s.validator.Validate(checkout); err != nil {
		return nil, customstatus.ErrBadRequest.Code, err
	}

	adminID := ctx.Locals("adminID").(string)
	email := ctx.Locals("email").(string)

	// create an errgroup.Group instance
	var g errgroup.Group

	tx, err := s.dbTransactionRepository.BeginTransaction()
	if err != nil {
		return nil, customstatus.ErrInternalServerError.Code, errors.New(customstatus.ErrInternalServerError.Message)
	}

	totalPrice := 0
	totalQuantity := 0
	for _, items := range input.Items {
		item := items
		g.Go(func() error {
			s.skuMutex.Lock()
			defer s.skuMutex.Unlock()

			// find the selected SKU and locking row (select for update)
			sku, err := s.skuRepository.GetSkuById(tx, item.ID)
			if err != nil {
				tx.Rollback()
				return errors.New(customstatus.ErrNotFound.Message)
			}
			if sku.Stock < item.Quantity {
				tx.Rollback()
				return fmt.Errorf("stock item %s tidak mencukupi hanya ada %d", sku.Name, sku.Stock)
			}

			sku.Stock -= item.Quantity

			err = s.skuRepository.UpdateStockSku(tx, &sku)
			if err != nil {
				tx.Rollback()
				return errors.New(customstatus.ErrInternalServerError.Message)
			}

			totalPrice += sku.Price * item.Quantity

			return nil
		})
		totalQuantity += item.Quantity
	}

	// wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		return nil, customstatus.ErrBadRequest.Code, errors.New(err.Error())
	}

	trx := &entity.Transaction{
		ID:               s.identifier.NewUuid(),
		AdminID:          adminID,
		TotalQuantity:    totalQuantity,
		TotalTransaction: totalPrice,
	}

	err = s.transactionRepository.Create(tx, trx)
	if err != nil {
		tx.Rollback()
		return nil, customstatus.ErrInternalServerError.Code, errors.New(customstatus.ErrInternalServerError.Message)
	}

	//store task to queue
	taskPayload := &redis.PayloadSendEmail{
		Email: email,
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),               //max retry when task failed execute
		asynq.ProcessIn(time.Second),     //sets the delay before processing a task to 1
		asynq.Queue(redis.QueueCritical), //it means this task will be prioritized.
	}

	codeError, err := s.workerTask.DistributeTaskSendEmail(taskPayload, opts...)
	if err != nil {
		tx.Rollback()
		return nil, codeError, errors.New(err.Error())
	}

	err = s.dbTransactionRepository.CommitTransaction(tx)
	if err != nil {
		tx.Rollback()
		return nil, customstatus.ErrInternalServerError.Code, errors.New(customstatus.ErrInternalServerError.Message)
	}

	return nil, customstatus.StatusCreated.Code, nil
}
