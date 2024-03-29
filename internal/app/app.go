package app

import (
	"gorm.io/gorm"
)

//go:generate mockgen -source app.go -package mocks -destination mocks/app_mock.go

type App interface {
	DB() *gorm.DB
	Pkg() *Pkg
	Repository() *Repository
	Service() *Service
	Log() Logger
}
