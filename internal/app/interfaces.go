package app

//go:generate mockgen -source interfaces.go -package app -destination mocks/interfaces_mock.go

type UserAccess interface {
}
