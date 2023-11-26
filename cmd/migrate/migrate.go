package main

import (
	"flag"
	"github.com/pressly/goose/v3"
	"go-template/internal/config"
	_ "go-template/migrations"
	"go-template/pkg/db/postgresdb"
	"log"
	"os"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "migrations", "directory with migration files")
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]

	gormdb, err := postgresdb.New(postgresdb.Config{
		Host:     config.Env.DatabaseHost,
		Port:     config.Env.DatabasePort,
		Username: config.Env.DatabaseUsername,
		Password: config.Env.DatabasePassword,
		Database: config.Env.DatabaseName,
		SSLMode:  config.Env.DatabaseSslMode,
	})
	if err != nil {
		panic(err)
	}

	db, err := gormdb.DB()
	if err != nil {
		log.Fatalf("gorm: failed to get *sql.DB connection: %v\n", err)
	}

	err = goose.SetDialect("postgres")
	if err != nil {
		log.Fatalf("goose: failed to set dialect: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 2 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
