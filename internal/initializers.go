package internal

import (
    "bets/config"
    "bets/internal/app"
    "bets/internal/repository"
    "bets/pkg/db/postgresdb"
    "bets/pkg/jwt"
    "bets/pkg/stub"
    "github.com/hibiken/asynq"
    "github.com/redis/go-redis/v9"
    "gorm.io/gorm"
)

func Startbetstion(appl *Application) error {
    if !appl.config.NoDB {
        db, err := InitializeDB()
        if err != nil {
            return err
        }
        *appl.db = *db
    } else {
        stubDb, err := stub.NewGorm()
        if err != nil {
            return err
        }
        *appl.db = *stubDb
    }

    ipkg, err := initializePkg(appl)
    if err != nil {
        return err
    }
    *appl.pkg = ipkg

    initializeRepository(appl)
    initializeService(appl)
    initializeTask(appl)

    return nil
}

func initializePkg(appl *Application) (app.Pkg, error) {
    // asynq task manager
    redisOpts, err := redis.ParseURL(config.Env.Fetch("REDIS_URL"))
    if err != nil {
        return app.Pkg{}, err
    }

    // redis client
    redisClient := redis.NewClient(redisOpts)

    var asynqClient app.AsynqClient = &stub.AsynqStub{}
    if !appl.config.NoRedis {
        asynqClient = asynq.NewClient(&asynq.RedisClientOpt{
            Network:      redisOpts.Network,
            Addr:         redisOpts.Addr,
            Username:     redisOpts.Username,
            Password:     redisOpts.Password,
            DB:           redisOpts.DB,
            DialTimeout:  redisOpts.DialTimeout,
            ReadTimeout:  redisOpts.ReadTimeout,
            WriteTimeout: redisOpts.WriteTimeout,
            PoolSize:     redisOpts.PoolSize,
            TLSConfig:    redisOpts.TLSConfig,
        })
    }

    return app.Pkg{
        AsynqClient: asynqClient,
        JWT:         jwt.New(config.Env.Fetch("JWT_SECRET_KEY")),
        RedisClient: redisClient,
    }, nil
}

func initializeRepository(appl *Application) {
    *appl.repository.Base.(*repository.Base) = *repository.NewBase(appl)
}

func initializeService(appl *Application) {
    //*appl.service.Auth.(*service.Auth) = *service.NewAuth(appl)
}

func initializeTask(appl *Application) {
    //*appl.task.SendEmail.(*task.SendEmail) = *task.NewSendEmail(appl)
}

func InitializeDB() (*gorm.DB, error) {
    db, err := postgresdb.New(postgresdb.Config{
        Url:                   config.Env.Fetch("DATABASE_URL"),
        SaveSQLAfterExecution: true,
    })

    return db, err
}
