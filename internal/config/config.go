package config

import (
	"fmt"
	"os"
	"strings"
)

var Env EnvVariables

func init() {
	// checking if we are running tests
	if strings.HasSuffix(os.Args[0], ".test") {
		if err := os.Setenv("GOLANG_ENV", "test"); err != nil {
			panic(err)
		}
	}

	Env = newEnv()
	err := Env.load()
	if err != nil {
		fmt.Println(err)
	}
}
