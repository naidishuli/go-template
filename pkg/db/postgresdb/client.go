package postgresdb

import (
    "fmt"
    "log"
    "os"
    "time"

    _ "github.com/lib/pq"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

func New(config Config) (*gorm.DB, error) {
    config.parse()

    psqlInfo := fmt.Sprintf("host=%s dbname=%s",
        config.Host,
        config.Database,
    )

    if config.Port > 0 {
        psqlInfo += fmt.Sprintf(" port=%d", config.Port)
    }

    if config.Username != "" {
        psqlInfo += fmt.Sprintf(" user=%s", config.Username)
    }

    if config.Password != "" {
        psqlInfo += fmt.Sprintf(" password=%s", config.Password)
    }

    if config.SSLMode == "" {
        config.SSLMode = "disable"

    }

    logConfig := logger.Config{
        SlowThreshold:             time.Second, // Slow SQL threshold
        LogLevel:                  logger.Info, // Log level
        IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
        ParameterizedQueries:      false,       // Don't include params in the SQL log
        Colorful:                  true,        // Disable color
    }

    if os.Getenv("GOLANG_ENV") == "production" {
        logConfig.LogLevel = logger.Silent
        logConfig.ParameterizedQueries = true
    }

    if os.Getenv("GOLANG_ENV") == "test" {
        logConfig.LogLevel = logger.Warn
        logConfig.ParameterizedQueries = true
    }

    logMode := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
        logConfig,
    )

    db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
        Logger:                 logMode,
        FullSaveAssociations:   false,
        SkipDefaultTransaction: true,
    })
    if err != nil {
        return nil, err
    }

    dbRef, err := db.DB()
    if err != nil {
        return nil, err
    }

    if err = dbRef.Ping(); err != nil {
        return nil, err
    }

    dbRef.SetMaxIdleConns(config.MaxIdleConnections)
    dbRef.SetMaxOpenConns(config.MaxOpenConnections)

    callbackFnc := dbStatementCallback(config)

    if config.SaveSQLAfterExecution {
        db.Callback().Query().Register("*", callbackFnc)
        db.Callback().Create().Register("*", callbackFnc)
        db.Callback().Update().Register("*", callbackFnc)
        db.Callback().Delete().Register("*", callbackFnc)
        db.Callback().Row().Register("*", callbackFnc)
        db.Callback().Raw().Register("*", callbackFnc)
    }

    return db, nil
}

func dbStatementCallback(config Config) func(db *gorm.DB) {
    return func(db *gorm.DB) {
        stmt := db.Statement
        sqlQuery := db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)
        db.Set(config.CallbackSqlKey, sqlQuery)
    }
}
