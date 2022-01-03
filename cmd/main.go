package main

import (
	"fmt"

	"go-template/internal/app"
)

func main() {
	appl, err := app.New()
	if err != nil {
		panic(err)
	}

	fmt.Println(appl)
}
