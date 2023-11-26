package tests

import (
	"database/sql"
	"fmt"
	"go-template/internal/config"
	"go-template/pkg/db/postgresdb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _db *gorm.DB
var _touchedTables map[string]bool

func GetDB(opts ...*sql.DB) *gorm.DB {
	if len(opts) > 0 {
		gormDB, err := gorm.Open(postgres.New(postgres.Config{
			Conn: opts[0],
		}), &gorm.Config{})

		if err != nil {
			panic(err)
		}

		return gormDB
	}

	if _db == nil {
		initializeDatabase()
	}
	return _db
}

func CleanDb() {
	for k, v := range _touchedTables {
		if k == "" {
			continue
		}

		if !v {
			continue
		}

		err := GetDB().Exec(fmt.Sprintf(`TRUNCATE TABLE %s CASCADE`, k)).Error
		if err != nil {
			panic(err)
		}

		_touchedTables[k] = false
	}
}

func CreateRecords(records ...interface{}) {
	db := GetDB()

	for _, r := range records {
		res := db.Session(&gorm.Session{FullSaveAssociations: true}).
			Model(r).
			Create(r)

		if res.Error != nil {
			fmt.Printf("failed to create %v\n\t %+v", r, res.Error)
			panic(res.Error)
		}
	}
}

func SaveRecords(records ...interface{}) {
	db := GetDB()

	for _, r := range records {
		res := db.Session(&gorm.Session{FullSaveAssociations: true}).
			Model(r).
			Save(r)

		if res.Error != nil {
			fmt.Printf("failed to save %v\n\t %+v", r, res.Error)
			panic(res.Error)
		}
	}
}

func UpdateRecords(records ...interface{}) {
	db := GetDB()

	for _, r := range records {
		res := db.Session(&gorm.Session{FullSaveAssociations: true}).
			Model(r).
			Updates(r)

		if res.Error != nil {
			fmt.Printf("failed to update %v\n\t %+v", r, res.Error)
			panic(res.Error)
		}
	}
}

func initializeDatabase() {
	_touchedTables = map[string]bool{}

	var err error
	_db, err = postgresdb.New(postgresdb.Config{
		Host:               "127.0.0.1",
		Port:               config.Env.DatabasePort,
		Username:           config.Env.DatabaseUsername,
		Password:           config.Env.DatabasePassword,
		Database:           "backend_test",
		MaxIdleConnections: config.Env.DatabaseMaxIdleConnections,
		MaxOpenConnections: config.Env.DatabaseMaxConnections,
		SSLMode:            config.Env.DatabaseSslMode,
	})

	_db.Callback().Query().Register("*", dbStatementCallback)
	_db.Callback().Create().Register("*", dbStatementCallback)
	_db.Callback().Update().Register("*", dbStatementCallback)
	_db.Callback().Delete().Register("*", dbStatementCallback)
	_db.Callback().Row().Register("*", dbStatementCallback)
	_db.Callback().Raw().Register("*", dbStatementCallback)

	if err != nil {
		panic(err)
	}
}

func dbStatementCallback(db *gorm.DB) {
	_touchedTables[db.Statement.Table] = true
}
