package service

type DepsContext interface {
	Deps() Deps
}

type Deps struct {
	Pkg        interface{}
	Repository interface{}
	Service    interface{}
}
