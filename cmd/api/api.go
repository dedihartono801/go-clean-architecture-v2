package api

import (
	"log"

	"github.com/dedihartono801/go-clean-architecture-v2/cmd/api/routes"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/queue/redis"
	handler "github.com/dedihartono801/go-clean-architecture-v2/internal/delivery/http"

	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/repository"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/admin"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/book"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/product"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/sku"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/transaction"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/user"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/identifier"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/validator"
	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Run(database *gorm.DB, taskDistributor redis.TaskDistributor) {

	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	dbTransactionRepository := repository.NewDbTransactionRepository(database)

	userRepository := repository.NewUserRepository(database)
	userService := user.NewUserService(userRepository, validator, identifier)
	userHandler := handler.NewUserHandler(userService)

	bookRepository := repository.NewBookRepository()
	bookService := book.NewService(bookRepository, validator)
	bookHandler := handler.NewBookHandler(bookService)

	adminRepository := repository.NewAdminRepository(database)
	adminService := admin.NewAdminService(adminRepository, validator, identifier)
	adminHandler := handler.NewAdminHandler(adminService)

	productService := product.NewProductService(validator, identifier)
	productHandler := handler.NewFilmHandler(productService)

	skuRepository := repository.NewSkuRepository(database)
	skuService := sku.NewSkuService(skuRepository, validator, identifier)
	skuHandler := handler.NewSkuHandler(skuService)

	transactionRepository := repository.NewTransactionRepository(database)
	transactionService := transaction.NewTransactionService(taskDistributor, dbTransactionRepository, transactionRepository, skuRepository, validator, identifier)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	app := fiber.New()

	routes.SetupRoutes(app)
	routes.AdminRouter(app, adminHandler)
	routes.UserRouter(app, userHandler)
	routes.BookRouter(app, bookHandler)
	routes.ProductRouter(app, productHandler)
	routes.SkuRouter(app, skuHandler)
	routes.TransactionRouter(app, transactionHandler)

	if err := app.Listen(":5001"); err != nil {
		log.Fatalf("listen: %s", err)
	}
}
