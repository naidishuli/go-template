package validator

import (
    "reflect"
    "strings"
    "time"

    "backend/utils"
    "github.com/go-playground/validator/v10"
)

func notZeroValue(fl validator.FieldLevel) bool {
    return !fl.Field().IsZero()
}

func timeframe(fl validator.FieldLevel) bool {
    var start, end time.Time
    var ok bool

    start, ok = fl.Field().Interface().(time.Time)
    if !ok {
        date := fl.Field().Interface().(utils.Date)
        start = date.Time
    }

    params := strings.Split(fl.Param(), " ")
    targetFieldValue := fl.Parent().FieldByName(params[0])

    if targetFieldValue.IsZero() {
        return true
    }

    if targetFieldValue.Kind() == reflect.Ptr {
        targetFieldValue = targetFieldValue.Elem()
    }

    end, ok = targetFieldValue.Interface().(time.Time)
    if !ok {
        date := targetFieldValue.Interface().(utils.Date)
        end = date.Time
    }

    if end.IsZero() {
        return true
    }

    return end.After(start)
}

func emptyIf(fl validator.FieldLevel) bool {
    params := strings.Split(fl.Param(), " ")
    targetFieldValue := fl.Parent().FieldByName(params[0])

    if params[1] == "empty" {
        return !targetFieldValue.IsZero()
    }

    return true
}
