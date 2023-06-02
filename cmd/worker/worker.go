package worker

import (
	"log"

	"github.com/dedihartono801/go-clean-architecture-v2/pkg/config"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/queue/redis"
	"github.com/hibiken/asynq"
)

func RunQueueRedisServer(config config.Config, redisOpt asynq.RedisClientOpt) {
	taskProcessor := redis.NewServer(redisOpt)
	err := taskProcessor.Start()
	if err != nil {
		log.Fatalf("failed to start worker")
	}
	log.Println("start worker")
}
