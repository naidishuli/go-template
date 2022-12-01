package app

type Repository struct {
	Temp TempRepo
}

type TempRepo interface {
	DoSomethingTemp(arg string, ctx *Context) error
}
