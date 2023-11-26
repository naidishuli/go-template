package config

type ContextType string

const (
	// gorm metadata
	GormSQLKey = "sql-query"

	// context metadata
	GormContextKey        ContextType = "gorm"
	GormClausesContextKey ContextType = "gorm-clauses"
	GormScopesContextKey  ContextType = "gorm-scopes"
	LoggerContextKey      ContextType = "logger"
	AsynqTaskContextKey   ContextType = "asynq-task"
	UserContextKey        ContextType = "user"
)
