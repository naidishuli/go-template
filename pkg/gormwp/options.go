package gormwp

import "gorm.io/gorm"

type Option func(db *gorm.DB) (*gorm.DB, error)

func Joins(query string, args ...interface{}) Option {
	return func(db *gorm.DB) (*gorm.DB, error) {
		ret := db.Joins(query, args)
		return ret, ret.Error
	}
}

func Model(value interface{}) Option {
	return func(db *gorm.DB) (*gorm.DB, error) {
		ret := db.Model(value)
		return ret, ret.Error
	}
}

func Preload(query string, args ...interface{}) Option {
	return func(db *gorm.DB) (*gorm.DB, error) {
		ret := db.Preload(query, args)
		return ret, ret.Error
	}
}

func Select(query interface{}, args ...interface{}) Option {
	return func(db *gorm.DB) (*gorm.DB, error) {
		ret := db.Select(query, args)
		return ret, ret.Error
	}
}

func Where(query interface{}, args ...interface{}) Option {
	return func(db *gorm.DB) (*gorm.DB, error) {
		ret := db.Where(query, args)
		return ret, ret.Error
	}
}

func Offset(offset int) Option {
	return func(db *gorm.DB) (*gorm.DB, error) {
		ret := db.Offset(offset)
		return ret, ret.Error
	}
}

func Limit(limit int) Option {
	return func(db *gorm.DB) (*gorm.DB, error) {
		ret := db.Limit(limit)
		return ret, ret.Error
	}
}
