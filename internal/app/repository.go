package app

//go:generate mockgen -source repository.go -package mocks -destination mocks/repository_mock.go

type Repository struct {
	Temp TempRepo
}

type TempRepo interface {
	DoSomethingTemp(ctx *Ctx, arg string) error
}
