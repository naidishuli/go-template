package main

import (
	"fmt"
	"go-template/api/routes"
	"time"

	"go-template/api"
	"go-template/internal"
	"go-template/internal/config"
)

func main() {
	location, err := time.LoadLocation("UTC")
	time.Local = location

	app, err := internal.NewApplication()
	if err != nil {
		panic(err)
	}

	fiberApp := api.New()

	routes.RegisterRoutes(app, fiberApp)
	err = fiberApp.Listen(fmt.Sprintf(":%d", config.Env.Port))
	if err != nil {
		panic(err)
	}
}
