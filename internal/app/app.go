package app

import "gorm.io/gorm"

type App interface {
    DB() *gorm.DB
    Log() Logger

    Pkg() Pkg
    Repository() Repository
    Service() Service
    Task() Task
    Cases() Cases
}
