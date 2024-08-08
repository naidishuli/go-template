package app

import (
    "context"

    "github.com/hibiken/asynq"
)

//go:generate mockgen -source task.go -package mocks -destination mocks/task_mock.go

type TaskHandler interface {
    Handle(ctx context.Context, t *asynq.Task) error
    Type() string
}

type Task struct {
}
