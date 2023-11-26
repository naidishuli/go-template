package main

import (
	"fmt"
	"time"

	"go-template/api"
	"go-template/internal"
	"go-template/internal/config"
)

func main() {
	location, err := time.LoadLocation("UTC")
	time.Local = location

	app := internal.NewApplication()
	err = internal.InitializeApp(app, internal.ApplicationConfig{})
	if err != nil {
		panic(err)
	}

	fiberApp := api.New()

	api.RegisterRoutes(app, fiberApp)
	err = fiberApp.Listen(fmt.Sprintf(":%d", config.Env.Port))
	if err != nil {
		panic(err)
	}
}
