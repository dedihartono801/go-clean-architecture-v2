package routes

import (
	"github.com/dedihartono801/go-clean-architecture-v2/cmd/api/middleware"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

// AdminRouter is the Router for GoFiber App
func AdminRouter(app fiber.Router, adminHandler http.AdminHandler) {
	adminRoute := app.Group("/admin")
	adminRoute.Post("/login", adminHandler.Login)
	adminRoute.Post("/create", adminHandler.Create)
	adminRoute.Get("", middleware.AuthUser, adminHandler.Find)
}
