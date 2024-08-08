package config

import (
    "embed"
    "fmt"
    "os"
    "time"

    "go-template/utils"
)

//go:embed *
var FS embed.FS

var Env environment

func init() {
    // checking if we are running tests
    // if strings.HasSuffix(os.Args[0], ".test") {
    //	if err := os.Setenv("GOLANG_ENV", "test"); err != nil {
    //		panic(err)
    //	}
    // }

    location, err := time.LoadLocation("UTC")
    time.Local = location
    if err != nil {
        panic(err)
    }

    err = loadEnvs()
    if err != nil {
        fmt.Println(err)
    }

    Env = environment{
        golangEnv: utils.FetchString(os.Getenv("GOLANG_ENV"), Development),
    }

    Env.load()
}
