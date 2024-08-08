package app

import (
    "go-template/pkg/rest"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

//go:generate mockgen -source repository.go -package mocks -destination mocks/repository_mock.go

type Repository struct {
    Base BaseRepo
}

type BaseRepoCtx interface {
    DB(defs ...*gorm.DB) *gorm.DB
}

type BaseRepo interface {
    WithCtx(ctx BaseRepoCtx) BaseRepo
    DB() *gorm.DB

    Clauses(conds ...clause.Expression) BaseRepo
    Joins(query string, args ...any) BaseRepo
    Model(value any) BaseRepo
    Select(query any, args ...any) BaseRepo
    Omit(columns ...string) BaseRepo
    Scopes(funcs ...func(*gorm.DB) *gorm.DB) BaseRepo
    Where(query any, args ...any) BaseRepo
    Unscoped() BaseRepo
    Transaction(fun func(repo BaseRepo) error) error
    TransactionCtx(ctx Context) Context

    Create(value any) error
    Delete(value any, cond ...any) error
    Find(dest any, cond ...any) error
    First(dest any, cond ...any) error
    FirstOrCreate(dest any, cond ...any) error
    Update(column string, value any) (int64, error)
    Updates(values any) error
    Save(value any) error

    Paginated(output any, args rest.PaginatedArgs) (int64, error)
}
