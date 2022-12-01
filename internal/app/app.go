package app

import "gorm.io/gorm"

type App interface {
	DB() *gorm.DB
	Pkg() *Pkg
	Repository() *Repository
	Service() *Service
}
