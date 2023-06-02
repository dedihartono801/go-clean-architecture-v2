package routes

import (
	"github.com/dedihartono801/go-clean-architecture-v2/cmd/api/middleware"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

// BookRouter is the Router for GoFiber App
func BookRouter(app fiber.Router, bookHandler http.BookHandler) {
	booksRoute := app.Group("/books", middleware.AuthUser)
	booksRoute.Get("", bookHandler.List)
	booksRoute.Get("/:id", bookHandler.Find)
	booksRoute.Put("/:id", bookHandler.Update)
	booksRoute.Post("", bookHandler.Create)
	booksRoute.Delete("/:id", bookHandler.Delete)
}
