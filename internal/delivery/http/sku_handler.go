package http

import (
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/sku"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/customstatus"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/dto"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

type SkuHandler interface {
	List(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}

type skuHandler struct {
	service sku.Service
}

func NewSkuHandler(service sku.Service) SkuHandler {
	return &skuHandler{service}
}

// List godoc
// @Summary      List sku
// @Tags         skus
// @Accept       json
// @Produce      json
// @Success      200  {array} entity.Sku
// @Security ApiKeyAuth
// @Router       /sku [get]
func (h *skuHandler) List(ctx *fiber.Ctx) error {
	skus, err := h.service.List()
	if err != nil {
		return helpers.CustomResponse(ctx, nil, err.Error(), customstatus.ErrInternalServerError.Code)
	}
	return helpers.CustomResponse(ctx, skus, customstatus.StatusOk.Message, customstatus.StatusOk.Code)
}

// Create godoc
// @Summary      Create sku
// @Tags         skus
// @Accept       json
// @Produce      json
// @Param		 raw	body	object		true	"body raw"
// @Success      200  {object} entity.Sku
// @Security ApiKeyAuth
// @Router       /sku [post]
func (h *skuHandler) Create(ctx *fiber.Ctx) error {
	skuDto := new(dto.SkuCreateDto)
	if err := ctx.BodyParser(skuDto); err != nil {
		return helpers.CustomResponse(ctx, nil, err.Error(), customstatus.ErrBadRequest.Code)
	}

	sku, statusCode, err := h.service.Create(skuDto)
	if err != nil {
		return helpers.CustomResponse(ctx, nil, err.Error(), statusCode)
	}
	return helpers.CustomResponse(ctx, sku, customstatus.StatusCreated.Message, statusCode)
}
