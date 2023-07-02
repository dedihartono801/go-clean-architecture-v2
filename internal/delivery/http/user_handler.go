package http

import (
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/http/user"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/customstatus"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/dto"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	List(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) UserHandler {
	return &userHandler{service}
}

// List godoc
// @Summary      List user
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array} entity.User
// @Security ApiKeyAuth
// @Router       /users [get]
func (h *userHandler) List(ctx *fiber.Ctx) error {
	users, err := h.service.List()
	if err != nil {
		return helpers.CustomResponse(ctx, nil, err.Error(), customstatus.ErrInternalServerError.Code)
	}
	return helpers.CustomResponse(ctx, users, customstatus.StatusOk.Message, customstatus.StatusOk.Code)
}

// Find godoc
// @Summary      Find user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Success      200  {object} entity.User
// @Security ApiKeyAuth
// @Router       /users/{id} [get]
func (h *userHandler) Find(ctx *fiber.Ctx) error {
	user, err := h.service.Find(ctx.Params("id"))
	if err != nil {
		return helpers.CustomResponse(ctx, nil, err.Error(), customstatus.ErrNotFound.Code)
	}
	return helpers.CustomResponse(ctx, user, customstatus.StatusOk.Message, customstatus.StatusOk.Code)
}

// Create godoc
// @Summary      Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param		 raw	body	object		true	"body raw"
// @Success      200  {object} entity.User
// @Security ApiKeyAuth
// @Router       /users [post]
func (h *userHandler) Create(ctx *fiber.Ctx) error {
	userDto := new(dto.UserCreateDto)
	if err := ctx.BodyParser(userDto); err != nil {
		return helpers.CustomResponse(ctx, nil, err.Error(), customstatus.ErrBadRequest.Code)
	}

	user, statusCode, err := h.service.Create(userDto)
	if err != nil {
		return helpers.CustomResponse(ctx, nil, err.Error(), statusCode)
	}
	return helpers.CustomResponse(ctx, user, customstatus.StatusCreated.Message, statusCode)
}

// Update godoc
// @Summary      Update user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Param		 raw	body	object	true	"body raw"
// @Success      200  {object} entity.User
// @Security ApiKeyAuth
// @Router       /users/{id} [put]
func (h *userHandler) Update(ctx *fiber.Ctx) error {
	userDto := new(dto.UserUpdateDto)
	if err := ctx.BodyParser(userDto); err != nil {
		return helpers.CustomResponse(ctx, nil, err.Error(), customstatus.ErrBadRequest.Code)
	}

	result, statusCode, err := h.service.Update(ctx.Params("id"), userDto)
	if err != nil {
		return helpers.CustomResponse(ctx, nil, err.Error(), statusCode)
	}
	return helpers.CustomResponse(ctx, result, customstatus.StatusCreated.Message, statusCode)

}

// Delete godoc
// @Summary      Delete user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Success      200
// @Security ApiKeyAuth
// @Router       /users/{id} [delete]
func (h *userHandler) Delete(ctx *fiber.Ctx) error {
	if err := h.service.Delete(ctx.Params("id")); err != nil {
		return helpers.CustomResponse(ctx, nil, err.Error(), customstatus.ErrInternalServerError.Code)
	}
	return helpers.CustomResponse(ctx, nil, customstatus.StatusOk.Message, customstatus.StatusOk.Code)

}
