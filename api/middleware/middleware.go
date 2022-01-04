package middleware

import "go-template/internal"

type Pkg interface {
}

type Repository interface {
}

type Service interface {
}

type Middleware struct {
	pkg     Pkg
	repo    Repository
	service Service
}

func New(app internal.AppContext) *Middleware {
	return &Middleware{
		pkg:     app.Pkg(),
		repo:    app.Repository(),
		service: app.Service(),
	}
}
