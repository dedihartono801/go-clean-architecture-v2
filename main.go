package main

import (
	"context"

	"github.com/dedihartono801/go-clean-architecture-v2/cmd/api"
	"github.com/dedihartono801/go-clean-architecture-v2/cmd/worker"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/config"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/database"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/queue/redis"
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

	mysql := database.InitMysql(envConfig)
	go worker.RunQueueRedisServer(*envConfig, redisOpt)
	api.Run(mysql, taskDistributor)
}
