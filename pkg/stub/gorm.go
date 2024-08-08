package stub

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm() (*gorm.DB, error) {
	db, _, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
