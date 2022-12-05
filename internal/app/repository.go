package app

//go:generate mockgen -source repository.go -package app -destination mocks/repository_mock.go

type DBName string

const (
	DefaultDB DBName = "default"
)

type Repository struct {
	Temp TempRepo
}

type TempRepo interface {
	DoSomethingTemp(arg string, ctx *Context) error
}
