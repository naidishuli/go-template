package app

type Service struct {
	Temp TempService
}

type TempService interface {
	DoSomething(arg string, ctx *Context) error
}
