package command

import (
    "fmt"
    "time"

    "github.com/gofiber/fiber/v3/middleware/static"
    "github.com/spf13/cobra"
    "go-template/config"
    "go-template/internal"
    "go-template/internal/api"
)

var RunServerCmd = &cobra.Command{
    Use:   "server",
    Short: "Run server",
    Run:   runServer,
}

func runServer(c *cobra.Command, args []string) {
    location, err := time.LoadLocation("UTC")
    time.Local = location

    app := internal.NewApplication(internal.ApplicationConfig{})
    err = internal.StartApplication(app)
    if err != nil {
        panic(err)
    }

    fiberApp := api.New()
    fiberApp.Get("/*", static.New("./public"))

    api.RegisterRoutes(app, fiberApp)
    err = fiberApp.Listen(fmt.Sprintf(":%s", config.Env.Fetch("PORT")))
    if err != nil {
        panic(err)
    }
}
