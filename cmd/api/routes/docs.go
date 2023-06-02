package routes

import (
	_ "github.com/dedihartono801/go-clean-architecture-v2/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
)

func SetupRoutes(
	app fiber.Router,

) {
	app.Get("/", monitor.New())
	app.Get("/docs/*", swagger.HandlerDefault)
}
