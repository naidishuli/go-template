package interfaces

import "go-template/internal/repository"

type Repository interface {
	Create(ctx *repository.Context, value interface{}, opts ...repository.QueryOption) error
	Count(ctx *repository.Context, count *int64, opts ...repository.QueryOption) error
	Find(ctx *repository.Context, out interface{}, opts ...repository.QueryOption) error
	First(ctx *repository.Context, out interface{}, opts ...repository.QueryOption) error
	Pluck(ctx *repository.Context, col string, out interface{}, opts ...repository.QueryOption) error
	Update(ctx *repository.Context, col string, val interface{}, opts ...repository.QueryOption) (int64, error)
	Updates(ctx *repository.Context, val interface{}, opts ...repository.QueryOption) (int64, error)
}
