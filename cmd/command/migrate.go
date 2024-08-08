package command

import (
    "context"
    "flag"
    "log"

    "bets/internal"
    _ "bets/migrations"
    "github.com/pressly/goose/v3"
    "github.com/spf13/cobra"
)

var RunMigrationCmd = &cobra.Command{
    Use:   "migration",
    Short: "Run migration",
    Run:   runMigration,
}

var (
    flags = flag.NewFlagSet("goose", flag.ExitOnError)
    dir   = flags.String("dir", "migrations", "directory with migration files")
)

func runMigration(cmd *cobra.Command, args []string) {
    if len(args) < 1 {
        flags.Usage()
        return
    }

    command := args[0]

    gormdb, err := internal.InitializeDB()
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
    if len(args) >= 2 {
        arguments = append(arguments, args[1:]...)
    }

    if err := goose.RunContext(context.Background(), command, db, *dir, arguments...); err != nil {
        log.Fatalf("goose %v: %v", command, err)
    }
}
