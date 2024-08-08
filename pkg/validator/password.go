package validator

import (
    "fmt"
    "reflect"
    "regexp"
    "unicode"

    "github.com/go-playground/validator/v10"
)

var operationalLogic = regexp.MustCompile(`^(\S+)\s(\S+)\s(\S+)\s(\S+)$`)

//
func passwordValidation(fl validator.FieldLevel) bool {
    //continueCheck, check := applyLogic(fl)
    //if !continueCheck || continueCheck && !check {
    //    return true
    //}

    if pwd, ok := fl.Field().Interface().(string); ok {
        var (
            hasMinLen = len(pwd) >= 8
            //hasUpper   = false
            //hasLower   = false
            hasNumber = false
            //hasSpecial = false
        )
        for _, c := range pwd {
            switch {
            case unicode.IsUpper(c):
                //hasUpper = true
            case unicode.IsLower(c):
                //hasLower = true
            case unicode.IsDigit(c):
                hasNumber = true
            case unicode.IsPunct(c) || unicode.IsSymbol(c):
                //hasSpecial = true
            }
        }
        return hasMinLen && hasNumber
    }
    return true // return true here if the field is not a string
}

// todo implement a more generic way for conditional check
func applyLogic(fl validator.FieldLevel) (bool, bool) {
    parts := operationalLogic.FindStringSubmatch(fl.Param())
    if len(parts) == 0 {
        return false, false
    }

    condition := parts[1]
    targetField := parts[2]
    operation := parts[3]
    value := parts[4]

    if condition == "if" {
        return ifLogic(fl, targetField, operation, value)
    }

    return false, false
}

func ifLogic(fl validator.FieldLevel, targetField, operation, value string) (bool, bool) {
    top := fl.Top()
    if top.Kind() == reflect.Ptr {
        top = top.Elem()
    }

    _, found := top.Type().FieldByName(targetField)
    if !found {
        return false, false
    }

    val := top.FieldByName(targetField)

    if val.Kind() == reflect.Ptr {
        if value == "nil" {
            return true, val.IsNil()
        }

        val = val.Elem()
    }

    if operation == "eq" {
        switch val.Kind() {
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            return true, fmt.Sprintf("%v", val.Interface()) == value

        case reflect.String:
            return true, val.String() == value
        case reflect.Bool:
            return true, fmt.Sprintf("%t", val.Interface()) == value
        }
    }

    return false, false
}
