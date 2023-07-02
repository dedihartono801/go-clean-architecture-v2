package http

import (
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/http/book"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/customstatus"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/dto"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

type BookHandler interface {
	List(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type bookHandler struct {
	service book.Service
}

func NewBookHandler(service book.Service) BookHandler {
	return &bookHandler{service: service}
}

// List godoc
// @Summary      List book
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {array} entity.Book
// @Security ApiKeyAuth
// @Router       /books [get]
func (b *bookHandler) List(ctx *fiber.Ctx) error {
	books, err := b.service.List()
	if err != nil {
		return helpers.CustomResponse(ctx, nil, customstatus.ErrInternalServerError.Message, customstatus.ErrInternalServerError.Code)
	}
	return helpers.CustomResponse(ctx, books, customstatus.StatusOk.Message, customstatus.StatusOk.Code)
}

// Find godoc
// @Summary      Find book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Success      200  {object} entity.Book
// @Security ApiKeyAuth
// @Router       /books/{id} [get]
func (b *bookHandler) Find(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	book, err := b.service.Find(id)
	if err != nil {
		return helpers.CustomResponse(ctx, nil, customstatus.ErrNotFound.Message, customstatus.ErrNotFound.Code)
	}
	return helpers.CustomResponse(ctx, book, customstatus.StatusOk.Message, customstatus.StatusOk.Code)
}

// Create godoc
// @Summary      Create book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param		 raw	body	object		true	"body raw"
// @Success      200  {object} entity.Book
// @Security ApiKeyAuth
// @Router       /books [post]
func (b *bookHandler) Create(ctx *fiber.Ctx) error {
	bookDto := new(dto.BookCreateDto)
	if err := ctx.BodyParser(bookDto); err != nil {
		return helpers.CustomResponse(ctx, nil, customstatus.ErrBadRequest.Message, customstatus.ErrBadRequest.Code)
	}

	book, err := b.service.Create(bookDto)
	if err != nil {
		return helpers.CustomResponse(ctx, nil, customstatus.ErrInternalServerError.Message, customstatus.ErrInternalServerError.Code)
	}
	return helpers.CustomResponse(ctx, book, customstatus.StatusCreated.Message, customstatus.StatusCreated.Code)
}

// Update godoc
// @Summary      Update book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Param		 raw	body	object	true	"body raw"
// @Success      200  {object} entity.Book
// @Security ApiKeyAuth
// @Router       /books/{id} [put]
func (b *bookHandler) Update(ctx *fiber.Ctx) error {
	bookDto := new(dto.BookUpdateDto)
	if err := ctx.BodyParser(bookDto); err != nil {
		return helpers.CustomResponse(ctx, nil, customstatus.ErrBadRequest.Message, customstatus.ErrBadRequest.Code)
	}

	id := ctx.Params("id")

	result, err := b.service.Update(id, bookDto)
	if err != nil {
		return helpers.CustomResponse(ctx, nil, customstatus.ErrInternalServerError.Message, customstatus.ErrInternalServerError.Code)
	}
	return helpers.CustomResponse(ctx, result, customstatus.StatusCreated.Message, customstatus.StatusCreated.Code)

}

// Delete godoc
// @Summary      Delete book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id		path	string		true	"ID"
// @Success      200
// @Security ApiKeyAuth
// @Router       /books/{id} [delete]
func (b *bookHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := b.service.Delete(id); err != nil {
		return helpers.CustomResponse(ctx, nil, customstatus.ErrInternalServerError.Message, customstatus.ErrInternalServerError.Code)
	}
	return helpers.CustomResponse(ctx, nil, customstatus.StatusOk.Message, customstatus.StatusOk.Code)

}
