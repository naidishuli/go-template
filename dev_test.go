package main

import (
	"fmt"
	"testing"
)

//func TestS(t *testing.T) {
//	notFound := app.NotFoundError(model.User{}, nil, nil)
//	check := app.IsErr(notFound, app.ErrCode(app.NotFoundErr))
//
//	fmt.Println(check)
//}
//
//func TestS3(t *testing.T) {
//	notFound := app.NotFoundError(model.User{}, nil, nil)
//	check := app.IsErrType(notFound, app.RepositoryErr(""))
//
//	fmt.Println(check)
//}

type Ctx struct {
	I *int
	S string
}

func TestS(t *testing.T) {
	i := 5
	fc := &Ctx{
		I: &i,
		S: "This is s",
	}

	var fc1 Ctx
	fc1 = *fc

	newI := 78

	fc1.I = &newI

	fmt.Println(*fc.I)
	fmt.Println(*fc1.I)
}
