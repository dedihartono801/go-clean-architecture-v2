package transaction

import (
	"errors"

	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/repository"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/customstatus"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/identifier"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/validator"
)

type Service interface {
	Find(id string) (*entity.Transaction, int, error)
}

type service struct {
	transactionRepository repository.TransactionRepository
	validator             validator.Validator
	identifier            identifier.Identifier
}

func NewGrpcTransactionService(
	transactionRepository repository.TransactionRepository,
	validator validator.Validator,
	identifier identifier.Identifier,
) Service {
	return &service{
		transactionRepository: transactionRepository,
		validator:             validator,
		identifier:            identifier,
	}
}

func (s service) Find(id string) (*entity.Transaction, int, error) {
	trx, err := s.transactionRepository.GetTrxById(id)
	if err != nil {
		return nil, customstatus.ErrNotFound.Code, errors.New(customstatus.ErrNotFound.Message)
	}
	return &trx, customstatus.StatusOk.Code, nil
}
