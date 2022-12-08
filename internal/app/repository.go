package app

//go:generate mockgen -source repository.go -package mocks -destination mocks/repository_mock.go

type DBName string

const (
	DefaultDB DBName = "default"
)

type Repository struct {
	Temp TempRepo
}

type TempRepo interface {
	DoSomethingTemp(ctx *Context, arg string) error
}
