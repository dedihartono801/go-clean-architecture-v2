package routes

import (
	"github.com/dedihartono801/go-clean-architecture-v2/cmd/http/middleware"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

// UserRouter is the Router for GoFiber App
func UserRouter(app fiber.Router, userHandler http.UserHandler) {
	usersRoute := app.Group("/users", middleware.AuthUser)
	usersRoute.Get("", userHandler.List)
	usersRoute.Get("/:id", userHandler.Find)
	usersRoute.Put("/:id", userHandler.Update)
	usersRoute.Post("", userHandler.Create)
	usersRoute.Delete("/:id", userHandler.Delete)
}
