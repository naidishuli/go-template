package api

import (
    "net/http"
    "time"

    "go-template/config"
    "go-template/internal/api/apierror"
    "go-template/internal/app"
    "github.com/gofiber/fiber/v3"
    "github.com/gofiber/fiber/v3/middleware/cors"
    "github.com/gofiber/fiber/v3/middleware/logger"
    "github.com/gofiber/fiber/v3/middleware/recover"
    "github.com/hibiken/asynq"
    "github.com/hibiken/asynqmon"
    "github.com/valyala/fasthttp/fasthttpadaptor"
)

func New() *fiber.App {
    fiberApp := fiber.New(fiber.Config{
        ErrorHandler: apierror.ErrorHandler,
    })

    fiberApp.Use(recover.New(recover.Config{
        EnableStackTrace: true,
    }))

    fiberApp.Use(cors.New(cors.Config{
        AllowOrigins:  []string{"*"},
        ExposeHeaders: []string{"*"},
    }))

    fiberApp.Use(logger.New(logger.Config{
        TimeFormat: time.RFC822,
        Format:     "[${time}] - ${ip}:${port} ${status} ${latency} - ${method} ${path}\n",
    }))

    return fiberApp
}

// RegisterRoutes used to register api routes to their handlers.
func RegisterRoutes(app app.App, fiberApp *fiber.App) {

    //cmnController := common.NewController(app)
    //cmnMiddlewares := common.NewMiddleware(app)

    // register all the secure routes here
    //securedAPI := fiberApp.Group("/api", cmnMiddlewares.Authorize)

    // initialize all controllers here

    pingRoute(fiberApp)

    // asynq task manager
    redisOpts, err := asynq.ParseRedisURI(config.Env.Fetch("REDIS_URL"))
    if err != nil {
        panic(err)
    }

    mon := asynqmon.New(asynqmon.Options{
        RootPath:     "/asynq",
        RedisConnOpt: redisOpts,
    })

    fiberApp.Group("/", WrapHandler(mon.ServeHTTP))
}

func pingRoute(app *fiber.App) {
    app.Get("/ping", func(ctx fiber.Ctx) error {
        return ctx.Status(200).SendString("ok")
    })
}

func WrapHandler(f func(http.ResponseWriter, *http.Request)) func(ctx fiber.Ctx) error {
    return func(ctx fiber.Ctx) error {
        fasthttpadaptor.NewFastHTTPHandler(http.HandlerFunc(f))(ctx.Context())
        return nil
    }
}
