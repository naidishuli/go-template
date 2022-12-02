package app

import "gorm.io/gorm"

// Context is passed through every layer of the application
// it will contain dynamic data like:
// db transactions for controlling through layers
// a context logger for individual request/call information
// user information
// etc...
type Context struct {
	TX *gorm.DB // database transaction to be used by the next layer
	// todo add logger
	// todo user info
	// todo a background golang context to add more dynamic data when needed
}

func NewContext(def *Context) *Context {
	if def != nil {
		return def
	}

	return &Context{}
}

func (c *Context) DB(def *gorm.DB) *gorm.DB {
	if c.TX == nil {
		return def
	}

	return c.TX
}
