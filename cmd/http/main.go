package main

import (
	"context"
	"log"

	"github.com/dedihartono801/go-clean-architecture-v2/cmd/http/routes"
	"github.com/dedihartono801/go-clean-architecture-v2/cmd/worker"
	"github.com/dedihartono801/go-clean-architecture-v2/database"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/queue/kafka"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/queue/redis"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/repository"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/http/admin"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/http/book"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/http/product"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/http/sku"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/http/transaction"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/http/user"
	handler "github.com/dedihartono801/go-clean-architecture-v2/internal/delivery/http"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/config"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/identifier"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/validator"
	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/hibiken/asynq"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @contact.name Dedi Hartono
// @contact.email dedihartono801@mail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	swagger.New(swagger.Config{
		Title:        "Swagger API",
		DeepLinking:  false,
		DocExpansion: "none",
	})

	envConfig := config.SetupEnvFile()
	redisOpt := asynq.RedisClientOpt{
		Addr: envConfig.RedisAddress,
	}
	ctx := context.Background()
	taskDistributor := redis.NewRedisTaskDistributor(redisOpt, ctx)
	kafkaProducer, err := kafka.NewKafkaProducer(envConfig.KafkaAdress)
	if err != nil {
		log.Fatalf(err.Error())
	}

	mysql := database.InitMysql(envConfig)

	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	dbTransactionRepository := repository.NewDbTransactionRepository(mysql)

	userRepository := repository.NewUserRepository(mysql)
	userService := user.NewUserService(userRepository, validator, identifier)
	userHandler := handler.NewUserHandler(userService)

	bookRepository := repository.NewBookRepository()
	bookService := book.NewService(bookRepository, validator)
	bookHandler := handler.NewBookHandler(bookService)

	adminRepository := repository.NewAdminRepository(mysql)
	adminService := admin.NewAdminService(adminRepository, validator, identifier)
	adminHandler := handler.NewAdminHandler(adminService)

	productService := product.NewProductService(validator, identifier)
	productHandler := handler.NewFilmHandler(productService)

	skuRepository := repository.NewSkuRepository(mysql)
	skuService := sku.NewSkuService(skuRepository, validator, identifier)
	skuHandler := handler.NewSkuHandler(skuService)

	transactionRepository := repository.NewTransactionRepository(mysql)
	transactionService := transaction.NewTransactionService(kafkaProducer, taskDistributor, dbTransactionRepository, transactionRepository, skuRepository, validator, identifier)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	app := fiber.New()

	routes.SetupRoutes(app)
	routes.AdminRouter(app, adminHandler)
	routes.UserRouter(app, userHandler)
	routes.BookRouter(app, bookHandler)
	routes.ProductRouter(app, productHandler)
	routes.SkuRouter(app, skuHandler)
	routes.TransactionRouter(app, transactionHandler)

	go worker.RunRedisWorker(*envConfig, redisOpt)
	go worker.RunKafkaWorker(envConfig.KafkaAdress, envConfig.ConsumerGroup, envConfig.Topic)

	if err := app.Listen(":5001"); err != nil {
		log.Fatalf("listen: %s", err)
	}
}
