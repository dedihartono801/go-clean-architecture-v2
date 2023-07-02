package http

import (
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/http/product"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/customstatus"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	Product(ctx *fiber.Ctx) error
}

type productHandler struct {
	service product.Service
}

func NewFilmHandler(service product.Service) ProductHandler {
	return &productHandler{service: service}
}

// Update godoc
// @Summary      List Product
// @Tags         product
// @Accept       json
// @Produce      json
// @Success      200  {object} dto.ProductDto
// @Security ApiKeyAuth
// @Router       /product [get]
func (h *productHandler) Product(ctx *fiber.Ctx) error {

	dt, statusCode, err := h.service.Product()
	if err != nil {
		return helpers.CustomResponse(ctx, nil, err.Error(), statusCode)
	}

	return helpers.CustomResponse(ctx, dt, customstatus.StatusOk.Message, statusCode)

}
