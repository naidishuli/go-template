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

type E struct {
	Message string
}

func (e *E) String() {
	fmt.Println(e.Message)
}

type EE E

func TestS(t *testing.T) {
	var m map[string]interface{}

	//a := m["dasd"]

	fmt.Println(m == nil)
}
