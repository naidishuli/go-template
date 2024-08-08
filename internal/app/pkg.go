package app

import (
    "context"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/hibiken/asynq"
    "github.com/redis/go-redis/v9"
)

//go:generate mockgen -source pkg.go -package mocks -destination mocks/pkg_mock.go

type Pkg struct {
    AsynqClient
    JWT
    RedisClient
}

type AsynqClient interface {
    Enqueue(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error)
    EnqueueContext(ctx context.Context, task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error)
    Close() error
}

type JWT interface {
    ParseToken(tokenValue string, claim jwt.Claims) error
    GenerateToken(claim jwt.Claims) (string, error)
}

type RedisClient interface {
    Get(ctx context.Context, key string) *redis.StringCmd
    Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}
