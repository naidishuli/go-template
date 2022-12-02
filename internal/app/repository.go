package app

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
