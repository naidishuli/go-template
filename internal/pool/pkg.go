package pool

// Pkg used as an initializer of the outside pool module
type Pkg struct {
}

func NewPkg() (Pkg, error) {
	return Pkg{}, nil
}
