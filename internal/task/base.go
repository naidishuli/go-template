package task

import (
    "encoding/json"

    "bets/internal/app"
    "github.com/hibiken/asynq"
)

type Base struct {
    client app.AsynqClient
    log    app.Logger

    typename string
    queue    string
    opts     []asynq.Option
}

func NewBase(dep app.App, typename string, queue string, opts ...asynq.Option) Base {
    baseOpts := []asynq.Option{asynq.Queue(queue)}
    if len(opts) > 0 {
        baseOpts = append(baseOpts, opts...)
    }

    return Base{
        client:   dep.Pkg().AsynqClient,
        log:      dep.Log(),
        typename: typename,
        queue:    queue,
        opts:     baseOpts,
    }
}

func (b Base) Enqueue(ctx *app.Ctx, data any, opts ...asynq.Option) (*asynq.TaskInfo, error) {
    payload, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    newOpts := b.opts
    for _, opt := range opts {
        if opt == nil {
            continue
        }

        newOpts = append(newOpts, opt)
    }

    task := asynq.NewTask(b.typename, payload, newOpts...)
    info, err := b.client.EnqueueContext(ctx, task)
    if err != nil {
        return nil, err
    }

    if info != nil {
        b.log.Infof("enqueued task : %s | type: %s | queue: %s", info.ID, info.Type, info.Queue)
    }

    return info, nil
}

func (b Base) Task(data any) *asynq.Task {
    var payload []byte
    var err error
    if data != nil {
        payload, err = json.Marshal(data)
        if err != nil {
            payload = nil
            b.log.Warnf("payload cannot be created: %+v", err)
        }
    }

    return asynq.NewTask(b.typename, payload, b.opts...)
}

func (b Base) Type() string {
    return b.typename
}
