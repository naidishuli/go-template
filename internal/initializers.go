package internal

import (
	"go-template/internal/app"
	"go-template/internal/config"
	"go-template/internal/service"
	"go-template/pkg/db/postgresdb"
	"gorm.io/gorm"
)

func Initialize(appl *Application, cfg ApplicationConfig) error {
	if !cfg.NoDB {
		db, err := initializeDB()
		if err != nil {
			return err
		}
		*appl.db = *db
	}

	pkgPool, err := initializePkg(appl, cfg)
	if err != nil {
		return err
	}

	*appl.pkgPool = pkgPool
	initializeRepository(appl)
	initializeService(appl)
	initializeTask(appl)

	return nil
}

func initializeRepository(appl *Application) {
	// *appl.repositoryPool.Base.(*repository.Base) = *repository.NewBase(appl)
}

func initializeService(appl *Application) {
	*appl.servicePool.Temp.(*service.Temp) = *service.NewTemp(appl)
}

func initializeTask(appl *Application) {

}

// todo add more context on errors here
func initializePkg(appl *Application, cfg ApplicationConfig) (app.Pkg, error) {
	// asynqClient := asynqwrp.NewClient(
	// 	asynq.RedisClientOpt{Addr: config.Env.RedisURL},
	// 	config.Env.SkipAsynq,
	// )

	if config.Env.GolangEnv == "production" {
		// todo use something else on the cloud
	}

	return app.Pkg{}, nil
}

func initializeDB() (*gorm.DB, error) {
	db, err := postgresdb.New(postgresdb.Config{
		Host:                  config.Env.DatabaseHost,
		Port:                  config.Env.DatabasePort,
		Username:              config.Env.DatabaseUsername,
		Password:              config.Env.DatabasePassword,
		Database:              config.Env.DatabaseName,
		SSLMode:               config.Env.DatabaseSslMode,
		SaveSQLAfterExecution: true,
	})

	return db, err
}
