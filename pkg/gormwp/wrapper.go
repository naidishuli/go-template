package gormwp

import "gorm.io/gorm"

type Context struct {
	TX *gorm.DB
}

type DB struct {
	db *gorm.DB
}

func New(db *gorm.DB) *DB {
	return &DB{
		db: db,
	}
}

func (d *DB) Create(value interface{}, opts []Option, ctx ...*Context) error {
	gdb := d.getDB(ctx)

	// Apply all the options
	var err error
	for _, opt := range opts {
		gdb, err = opt(gdb)
		if err != nil {
			return err
		}
	}

	return gdb.Create(value).Error
}

func (d *DB) Count(count *int64, opts []Option, ctx ...*Context) error {
	gdb := d.getDB(ctx)

	// Apply all the options
	var err error
	for _, opt := range opts {
		gdb, err = opt(gdb)
		if err != nil {
			return err
		}
	}

	return gdb.Count(count).Error
}

func (d *DB) Find(out interface{}, opts []Option, ctx ...*Context) error {
	gdb := d.getDB(ctx)

	// Apply all the options
	var err error
	for _, opt := range opts {
		gdb, err = opt(gdb)
		if err != nil {
			return err
		}
	}

	return gdb.Find(out).Error
}

func (d *DB) First(out interface{}, opts []Option, ctx ...*Context) error {
	gdb := d.getDB(ctx)

	// Apply all the options
	var err error
	for _, opt := range opts {
		gdb, err = opt(gdb)
		if err != nil {
			return err
		}
	}

	return gdb.First(out).Error
}

func (d *DB) Pluck(col string, out interface{}, opts []Option, ctx ...*Context) error {
	gdb := d.getDB(ctx)

	// Apply all the options
	var err error
	for _, opt := range opts {
		gdb, err = opt(gdb)
		if err != nil {
			return err
		}
	}

	return gdb.Pluck(col, out).Error
}

func (d *DB) Update(col string, val interface{}, opts []Option, ctx ...*Context) (int64, error) {
	gdb := d.getDB(ctx)

	// Apply all the options
	var err error
	for _, opt := range opts {
		gdb, err = opt(gdb)
		if err != nil {
			return 0, err
		}
	}

	res := gdb.Update(col, val)
	return res.RowsAffected, res.Error
}

func (d *DB) Updates(val interface{}, opts []Option, ctx ...*Context) (int64, error) {
	gdb := d.getDB(ctx)

	// Apply all the options
	var err error
	for _, opt := range opts {
		gdb, err = opt(gdb)
		if err != nil {
			return 0, err
		}
	}

	res := gdb.Updates(val)
	return res.RowsAffected, res.Error
}

func (d *DB) Transaction(f func(tx *gorm.DB) error) error {
	if d.db == nil {
		return f(nil)
	}

	return d.Transaction(f)
}

func (d *DB) getDB(ctx []*Context) *gorm.DB {
	if len(ctx) == 0 || ctx[0].TX == nil {
		return d.db
	}

	return ctx[0].TX
}
