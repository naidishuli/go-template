package app

//go:generate mockgen -source service.go -package mocks -destination mocks/service_mock.go

type Service struct {
	Temp TempService
}

type TempService interface {
	DoSomething(ctx *Ctx, arg string) error
}
