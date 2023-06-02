package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"log"

	"github.com/hibiken/asynq"
	"github.com/mailgun/mailgun-go/v4"
)

const (
	TaskSendEmail = "task:send_email"
	QueueCritical = "critical"
	QueueDefault  = "default"
)

type PayloadSendEmail struct {
	Email string `json:"email"`
}

type TaskServer interface {
	Start() error
	ProcessTaskSendEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskServer struct {
	server *asynq.Server
}

// Create and configuring Asynq worker server.
func NewServer(redisOpt asynq.RedisClientOpt) TaskServer {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			// Specify how many concurrent workers to use.
			Concurrency: 10,
			// Specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6, // processed 60% of the time
				"default":  3, // processed 30% of the time
				"low":      1, // processed 10% of the time
			},
		})

	return &RedisTaskServer{
		server: server,
	}
}

func (processor *RedisTaskServer) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendEmail, processor.ProcessTaskSendEmail)

	return processor.server.Start(mux)
}

func (processor *RedisTaskServer) ProcessTaskSendEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_KEY"))
	message := mg.NewMessage(
		"noreply@gmail.com",
		"info pesanan",
		"Pesanan anda sudah kami terima",
		payload.Email,
	)

	_, _, err := mg.Send(ctx, message)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email sent successfully!")
	}
	return nil
}
