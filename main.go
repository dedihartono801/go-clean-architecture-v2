package main

import (
	"context"
	"log"

	"github.com/dedihartono801/go-clean-architecture-v2/cmd/api"
	"github.com/dedihartono801/go-clean-architecture-v2/cmd/worker"
	"github.com/dedihartono801/go-clean-architecture-v2/database"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/queue/kafka"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/queue/redis"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/config"
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
	go worker.RunRedisWorker(*envConfig, redisOpt)
	go worker.RunKafkaWorker(envConfig.KafkaAdress, envConfig.ConsumerGroup, envConfig.Topic)
	api.Run(mysql, taskDistributor, kafkaProducer)
}
