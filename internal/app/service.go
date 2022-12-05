package app

//go:generate mockgen -source service.go -package app -destination mocks/service_mock.go

type Service struct {
	Temp TempService
}

type TempService interface {
	DoSomething(arg string, ctx *Context) error
}
