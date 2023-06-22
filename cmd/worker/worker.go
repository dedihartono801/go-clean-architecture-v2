package worker

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/queue/kafka"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/queue/redis"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/config"
	"github.com/hibiken/asynq"
)

func RunRedisWorker(config config.Config, redisOpt asynq.RedisClientOpt) {
	taskProcessor := redis.NewServer(redisOpt)
	err := taskProcessor.Start()
	if err != nil {
		log.Fatalf("failed to start worker")
	}
	log.Println("start worker")
}

func RunKafkaWorker(broker string, groupConsumer string, topic string) {

	ready := make(chan bool)

	kafkaConsumer, err := kafka.NewKafkaConsumer(broker, groupConsumer)
	if err != nil {
		log.Fatal(err)
	}

	err = kafkaConsumer.StartConsumerGroup(ready, topic)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for the consumer group to be ready
	<-ready
	fmt.Println("Consumer group is ready")

	// Set up a signal handler to gracefully handle termination signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Wait for a termination signal
	<-signals
	fmt.Println("Terminating...")

	// Signal the consumer group to stop gracefully
	kafkaConsumer.CloseConsumerGroup()

	// Wait for the consumer group to finish before exiting
	<-ready
	fmt.Println("Consumer group terminated")
}
