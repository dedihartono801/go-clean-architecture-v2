package http

import (
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/transaction"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/customstatus"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/dto"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler interface {
	Checkout(ctx *fiber.Ctx) error
}

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) TransactionHandler {
	return &transactionHandler{service}
}

// Create godoc
// @Summary      Checkout Items
// @Tags         checkout
// @Accept       json
// @Produce      json
// @Param		 raw	body	object		true	"body raw"
// @Success      201  {object} entity.Transaction
// @Security ApiKeyAuth
// @Router       /checkout [post]
func (h *transactionHandler) Checkout(ctx *fiber.Ctx) error {
	checkoutDto := new(dto.TransactionCheckoutDto)
	if err := ctx.BodyParser(checkoutDto); err != nil {
		return helpers.CustomResponse(ctx, nil, err.Error(), customstatus.ErrBadRequest.Code)
	}

	transaction, statusCode, err := h.service.Checkout(ctx, checkoutDto)
	if err != nil {
		return helpers.CustomResponse(ctx, nil, err.Error(), statusCode)
	}
	return helpers.CustomResponse(ctx, transaction, customstatus.StatusCreated.Message, statusCode)
}
