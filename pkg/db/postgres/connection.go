package postgres

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbUrlRegex = regexp.MustCompile("//(.+):(.+)@(.+):(.+)/(.+)")

type Config struct {
	Url                string
	Host               string
	Port               int
	Username           string
	Password           string
	Database           string
	SSLMode            string
	MaxIdleConnections int
	MaxOpenConnections int
}

func (c *Config) parse() {
	if c.SSLMode == "" {
		c.SSLMode = "disable"
	}

	if c.MaxOpenConnections == 0 {
		c.MaxOpenConnections = 50
	}

	if c.MaxIdleConnections == 0 {
		c.MaxIdleConnections = 50
	}

	if c.Url != "" {
		parts := dbUrlRegex.FindStringSubmatch(c.Url)
		c.Username = parts[1]
		c.Password = parts[2]
		c.Host = parts[3]
		c.Port, _ = strconv.Atoi(parts[4])
		c.Database = parts[5]
	}
}

func NewPostgresConn(config Config) (*gorm.DB, error) {
	config.parse()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.Database,
		config.SSLMode,
	)

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

	return db, nil
}
