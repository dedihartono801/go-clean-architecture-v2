package routes

import (
	"github.com/dedihartono801/go-clean-architecture-v2/cmd/api/middleware"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

// ProductRouter is the Router for GoFiber App
func ProductRouter(app fiber.Router, productHandler http.ProductHandler) {
	productRoute := app.Group("/product", middleware.AuthUser)
	productRoute.Get("", productHandler.Product)
}
