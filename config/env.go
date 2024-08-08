package config

import (
    "os"
    "path"

    "bets/utils"
    "github.com/joho/godotenv"
)

type Specifics struct {
    TransactionalEmailFrom string
}

type environment struct {
    golangEnv string

    specifics Specifics
}

func (e *environment) Fetch(key string, args ...string) string {
    var def string
    if len(args) > 0 {
        def = args[0]
    }

    return utils.FetchString(os.Getenv(key), def)
}

func (e *environment) Specifics() Specifics {
    return e.specifics
}

func (e *environment) load() {
    switch e.golangEnv {
    case Development:
        e.loadDevelopment()
    case Testing:
        e.loadTesting()
    case Staging:
        e.loadStaging()
    case Production:
        e.loadProduction()
    }
}

func (e *environment) loadDevelopment() {
    e.specifics = Specifics{
        TransactionalEmailFrom: "naidishuli@gmail.com",
    }
}

// todo add right rules
func (e *environment) loadTesting() {
    e.specifics = Specifics{
        TransactionalEmailFrom: "noreply@test.com",
    }
}

// todo add right rules
func (e *environment) loadStaging() {
    e.specifics = Specifics{
        TransactionalEmailFrom: "staging-noreply@jobify.so",
    }
}

// todo add right rules
func (e *environment) loadProduction() {
    e.specifics = Specifics{
        TransactionalEmailFrom: "noreply@jobify.so",
    }
}

func (e *environment) IsDevelopment() bool {
    return e.golangEnv == Development
}

func (e *environment) IsTesting() bool {
    return e.golangEnv == Testing
}

func (e *environment) IsStaging() bool {
    return e.golangEnv == Staging
}

func (e *environment) IsProduction() bool {
    return e.golangEnv == Production
}

func loadEnvs() (err error) {
    switch os.Getenv("GOLANG_ENV") {
    case "test":
        err = loadEnvFile(".env.test")
    case "development":
        err = loadEnvFile(".env")
    case "production":

    default:
        err = loadEnvFile(".env")
    }

    return
}

func loadEnvFile(filename string) error {
    if filename == "" {
        return nil
    }

    envPath := path.Join(
        utils.RootPath(),
        filename,
    )

    return godotenv.Load(envPath)
}
