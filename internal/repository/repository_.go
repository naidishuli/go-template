package repository

import "gorm.io/gorm"

type Repository interface {
	Create(ctx *Context, value interface{}, opts ...QueryOption) error
	Count(ctx *Context, count *int64, opts ...QueryOption) error
	Find(ctx *Context, out interface{}, opts ...QueryOption) error
	First(ctx *Context, out interface{}, opts ...QueryOption) error
	Pluck(ctx *Context, col string, out interface{}, opts ...QueryOption) error
	Update(ctx *Context, col string, val interface{}, opts ...QueryOption) (int64, error)
	Updates(ctx *Context, val interface{}, opts ...QueryOption) (int64, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx *Context, value interface{}, opts ...QueryOption) error {
	c := r.getContext(ctx)
	gdb := c.Db

	var err error = nil
	// Apply all the options
	for _, opt := range opts {
		gdb, err = opt(gdb)
		if err != nil {
			return err
		}
	}

	return gdb.Create(value).Error
}

func (r *repository) Count(ctx *Context, count *int64, opts ...QueryOption) error {
	c := r.getContext(ctx)
	gdb := c.Db
	var err error = nil

	// Apply all the options
	for _, opt := range opts {
		gdb, err = opt(gdb)
		if err != nil {
			return err
		}
	}

	return gdb.Count(count).Error
}

func (r *repository) Find(ctx *Context, out interface{}, opts ...QueryOption) error {
	c := r.getContext(ctx)
	gdb := c.Db
	var err error = nil

	// Apply all the options
	for _, opt := range opts {
		gdb, err = opt(gdb)
		if err != nil {
			return err
		}
	}

	return gdb.Find(out).Error
}

func (r *repository) First(ctx *Context, out interface{}, opts ...QueryOption) error {
	c := r.getContext(ctx)
	gdb := c.Db
	var err error = nil

	// Apply all the options
	for _, opt := range opts {
		gdb, err = opt(gdb)
		if err != nil {
			return err
		}
	}

	return gdb.First(out).Error
}

func (r *repository) Pluck(ctx *Context, col string, out interface{}, opts ...QueryOption) error {
	c := r.getContext(ctx)
	gdb := c.Db
	var err error = nil

	// Apply all the options
	for _, opt := range opts {
		gdb, err = opt(gdb)
		if err != nil {
			return err
		}
	}

	return gdb.Pluck(col, out).Error
}

func (r *repository) Update(ctx *Context, col string, val interface{}, opts ...QueryOption) (int64, error) {
	c := r.getContext(ctx)
	gdb := c.Db
	var err error = nil

	// Apply all the options
	for _, opt := range opts {
		gdb, err = opt(gdb)
		if err != nil {
			return 0, err
		}
	}

	res := gdb.Update(col, val)
	return res.RowsAffected, res.Error
}

func (r *repository) Updates(ctx *Context, val interface{}, opts ...QueryOption) (int64, error) {
	c := r.getContext(ctx)
	gdb := c.Db
	var err error = nil

	// Apply all the options
	for _, opt := range opts {
		gdb, err = opt(gdb)
		if err != nil {
			return 0, err
		}
	}

	res := gdb.Updates(val)
	return res.RowsAffected, res.Error
}

func (r *repository) getContext(ctxs *Context) *Context {
	if ctxs == nil {
		return &Context{
			Db: r.db,
		}
	}

	ctx := ctxs

	if ctx.Db == nil {
		ctx.Db = r.db
	}

	if ctx.Test {
		ctx.Db = ctx.Db.Session(&gorm.Session{DryRun: true})
		ctx.Db.Callback().Query().Register("*", ctx.statementCallback)
		ctx.Db.Callback().Create().Register("*", ctx.statementCallback)
		ctx.Db.Callback().Update().Register("*", ctx.statementCallback)
		ctx.Db.Callback().Delete().Register("*", ctx.statementCallback)
		ctx.Db.Callback().Row().Register("*", ctx.statementCallback)
		ctx.Db.Callback().Raw().Register("*", ctx.statementCallback)
	}

	return ctx
}

type Context struct {
	Db       *gorm.DB
	Preloads []string
	Joins    []string
	Test     bool
	SQLs     []string
	Vars     [][]interface{}
}

func (c *Context) addRelations(query *gorm.DB) *gorm.DB {
	for _, join := range c.Joins {
		query = query.Joins(join)
	}

	for _, preload := range c.Preloads {
		query = query.Preload(preload)
	}

	return query
}

func (c *Context) statementCallback(db *gorm.DB) {
	stmt := db.Statement
	c.SQLs = append(c.SQLs, stmt.SQL.String())
	c.Vars = append(c.Vars, stmt.Vars)
}
