package routes

import (
	"github.com/dedihartono801/go-clean-architecture-v2/cmd/http/middleware"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

// SkuRouter is the Router for GoFiber App
func SkuRouter(app fiber.Router, skuHandler http.SkuHandler) {
	skuRoute := app.Group("/sku", middleware.AuthUser)
	skuRoute.Get("", skuHandler.List)
	skuRoute.Post("", skuHandler.Create)
}
