package postgresdb

import (
	"fmt"
	_ "github.com/lib/pq"
	"go-template/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func New(config Config) (*gorm.DB, error) {
	config.parse()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.Database,
	)

	if config.SSLMode != "" {
		psqlInfo += fmt.Sprintf(" sslmode=%s", config.SSLMode)
	}

	logMode := logger.Default.LogMode(logger.Info)
	if os.Getenv("GOLANG_ENV") == "production" {
		logMode = logger.Default.LogMode(logger.Silent)
	}

	if os.Getenv("GOLANG_ENV") == "test" {
		logMode = logger.Default.LogMode(logger.Warn)
	}

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

	if config.SaveSQLAfterExecution {
		db.Callback().Query().Register("*", dbStatementCallback)
		db.Callback().Create().Register("*", dbStatementCallback)
		db.Callback().Update().Register("*", dbStatementCallback)
		db.Callback().Delete().Register("*", dbStatementCallback)
		db.Callback().Row().Register("*", dbStatementCallback)
		db.Callback().Raw().Register("*", dbStatementCallback)
	}

	return db, nil
}

func dbStatementCallback(db *gorm.DB) {
	stmt := db.Statement
	sqlQuery := db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)
	db.Set(config.GormSQLKey, sqlQuery)
}
