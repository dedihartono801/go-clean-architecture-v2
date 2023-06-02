package http

import (
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/admin"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/customstatus"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/dto"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

type AdminHandler interface {
	Login(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}

type adminHandler struct {
	service admin.Service
}

func NewAdminHandler(service admin.Service) AdminHandler {
	return &adminHandler{service}
}

// Find godoc
// @Summary      Get Profile Admin
// @Tags         admin
// @Accept       json
// @Produce      json
// @Success      200  {object} entity.Admin
// @Security ApiKeyAuth
// @Router       /admin [get]
func (h *adminHandler) Find(ctx *fiber.Ctx) error {
	admin, err := h.service.Find(ctx.Locals("adminID").(string))
	if err != nil {
		return helpers.CustomResponse(ctx, nil, customstatus.ErrNotFound.Message, customstatus.ErrNotFound.Code)
	}
	return helpers.CustomResponse(ctx, admin, customstatus.StatusOk.Message, customstatus.StatusOk.Code)
}

// Create godoc
// @Summary      Create Admin
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param		 raw	body	object		true	"body raw"
// @Success      200  {object} entity.Admin
// @Router       /admin/create [post]
func (h *adminHandler) Create(ctx *fiber.Ctx) error {
	adminDto := new(dto.AdminCreateDto)
	if err := ctx.BodyParser(adminDto); err != nil {
		return helpers.CustomResponse(ctx, nil, err, 500)
	}

	admin, statusCode, err := h.service.Create(adminDto)
	if err != nil {
		return helpers.CustomResponse(ctx, nil, err.Error(), statusCode)
	}
	return helpers.CustomResponse(ctx, admin, customstatus.StatusCreated.Message, customstatus.StatusCreated.Code)
}

// Update godoc
// @Summary      Login admin
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param		 raw	body	object	true	"body raw"
// @Success      200  {object} entity.Admin
// @Router       /admin/login [post]
func (h *adminHandler) Login(ctx *fiber.Ctx) error {
	adminDto := new(dto.AdminLoginDto)
	if err := ctx.BodyParser(adminDto); err != nil {
		return err
	}

	result, statusCode, err := h.service.Login(adminDto)
	if err != nil {
		return helpers.CustomResponse(ctx, nil, err.Error(), statusCode)
	}

	return helpers.CustomResponse(ctx, result, customstatus.StatusOk.Message, customstatus.StatusOk.Code)

}
