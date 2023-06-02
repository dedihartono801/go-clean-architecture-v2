package redis

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/dedihartono801/go-clean-architecture-v2/pkg/customstatus"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

type TaskDistributor interface {
	DistributeTaskSendEmail(
		payload *PayloadSendEmail,
		opts ...asynq.Option,
	) (int, error)
}

type RedisTaskDistributor struct {
	client *asynq.Client
	ctx    context.Context
}

func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt, ctx context.Context) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributor{
		client: client,
		ctx:    ctx,
	}
}

func (distributor *RedisTaskDistributor) DistributeTaskSendEmail(
	payload *PayloadSendEmail,
	opts ...asynq.Option,
) (int, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return customstatus.ErrBadRequest.Code, errors.New(err.Error())
	}

	task := asynq.NewTask(TaskSendEmail, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(distributor.ctx, task)
	if err != nil {
		return customstatus.ErrInternalServerError.Code, errors.New(customstatus.ErrInternalServerError.Message)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return customstatus.ErrInternalServerError.Code, nil
}
