package app

//go:generate mockgen -source cases.go -package mocks -destination mocks/cases_mock.go

type Cases interface {
}
