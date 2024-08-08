package stub

import (
	"context"
	"github.com/hibiken/asynq"
)

type AsynqStub struct {
}

func (a AsynqStub) Enqueue(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return &asynq.TaskInfo{}, nil
}

func (a AsynqStub) EnqueueContext(ctx context.Context, task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return &asynq.TaskInfo{}, nil
}

func (a AsynqStub) Close() error {
	return nil
}
