package config

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"

	"github.com/joho/godotenv"

	"go-template/utils"
)

func newEnv() EnvVariables {
	return EnvVariables{
		loadEnvFile: loadEnvFileI,
		loadValues:  loadValuesI,
	}
}

type EnvVariables struct {
	GolangEnv                  string
	Port                       int
	DatabaseUrl                string
	DatabaseHostname           string
	DatabasePort               int
	DatabaseDbName             string
	DatabaseUsername           string
	DatabasePassword           string
	DatabaseSslMode            string
	DatabaseMaxIdleConnections int
	DatabaseMaxConnections     int
	JwtVerificationKey         string
	loadEnvFile                func(filename string) error
	loadValues                 func(*EnvVariables)
}

func (e *EnvVariables) IsProduction() bool {
	return "production" == os.Getenv("GOLANG_ENV")
}

func (e *EnvVariables) IsDevelopment() bool {
	return "development" == os.Getenv("GOLANG_ENV")
}

func (e *EnvVariables) IsTesting() bool {
	return "test" == os.Getenv("GOLANG_ENV")
}

func (e *EnvVariables) load() (err error) {
	switch os.Getenv("GOLANG_ENV") {
	case "test":
		err = e.loadEnvFile(".env.test")
	case "development":
		err = e.loadEnvFile(".env")
	case "production":

	default:
		err = e.loadEnvFile(".env")
	}

	// err = e.loadEnvFile(".env")
	e.loadValues(e)

	return
}

var loadEnvFileI = func(filename string) (err error) {
	if filename == "" {
		return fmt.Errorf("empty filename, specify the file that holds the environment variables")
	}

	envPath := path.Join(
		utils.RootPath(),
		filename,
	)

	err = godotenv.Load(envPath)
	return
}

var loadValuesI = func(e *EnvVariables) {
	values := reflect.ValueOf(e).Elem()
	fields := reflect.TypeOf(e).Elem()

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)
		snakeCaseName := strings.ToUpper(utils.ToSnakeCase(field.Name))

		envValue := os.Getenv(field.Name)
		if envValue == "" {
			envValue = os.Getenv(snakeCaseName)
		}

		switch value.Kind() {
		case reflect.Int:
			v, _ := strconv.ParseInt(envValue, 10, 0)
			value.SetInt(v)
		case reflect.String:
			value.SetString(envValue)
		}
	}
}
