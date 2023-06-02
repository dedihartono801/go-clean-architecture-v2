package routes

import (
	"github.com/dedihartono801/go-clean-architecture-v2/cmd/api/middleware"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

// TransactionRouter is the Router for GoFiber App
func TransactionRouter(app fiber.Router, transactionHandler http.TransactionHandler) {
	app.Post("/checkout", middleware.AuthUser, transactionHandler.Checkout)
}
