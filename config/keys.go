package config

type ContextType string

const (
	Development = "development"
	Staging     = "staging"
	Production  = "production"
	Testing     = "testing"

	// gorm metadata
	GormSQLKey = "sql-query"

	// context metadata
	GormContextKey        ContextType = "gorm"
	GormClausesContextKey ContextType = "gorm-clauses"
	GormScopesContextKey  ContextType = "gorm-scopes"
	GormSessionContextKey ContextType = "gorm-session"
	LoggerContextKey      ContextType = "logger"
	AsynqTaskContextKey   ContextType = "asynq-task"
	UserContextKey        ContextType = "user"

	GoogleSignInState = "google-sign-in"
	GoogleSignUpState = "google-sign-up"
)
