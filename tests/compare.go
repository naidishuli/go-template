package tests

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go-template/internal/app"
)

func DiffErrors(expected, actual any) string {
	_, ok := expected.(*app.Error)
	if ok {
		return cmp.Diff(expected, actual, cmpopts.IgnoreUnexported(app.Error{}))
	}

	return cmp.Diff(expected, actual, cmpopts.EquateErrors())
}
