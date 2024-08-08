package validator

import (
    "encoding/json"
    "fmt"
    "reflect"
    "strings"

    "github.com/go-playground/validator/v10"
)

type Wrapper struct {
    *validator.Validate
}

func (v *Wrapper) Struct(s any) error {
    err := v.Validate.Struct(s)
    if err != nil {
        return Map(err)
    }

    return nil
}

var instance *Wrapper

func Get() *Wrapper {
    if instance != nil {
        return instance
    }

    instance = &Wrapper{validator.New(validator.WithRequiredStructEnabled())}

    instance.RegisterTagNameFunc(func(fld reflect.StructField) string {
        name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

        if name == "-" {
            return ""
        }

        return name
    })

    err := instance.RegisterValidation("dateOfBirth", dataOfBirth)
    if err != nil {
        panic(err)
    }

    err = instance.RegisterValidation("emptyIf", emptyIf)
    if err != nil {
        panic(err)
    }

    err = instance.RegisterValidation("notZeroValue", notZeroValue)
    if err != nil {
        panic(err)
    }

    err = instance.RegisterValidation("password", passwordValidation)
    if err != nil {
        panic(err)
    }

    err = instance.RegisterValidation("timeframe", timeframe)
    if err != nil {
        panic(err)
    }

    return instance
}

// TODO add better errors messages for tags by using translations

func Map(err error) Errors {
    if err == nil {
        return nil
    }

    errors := make(Errors)
    for _, err := range err.(validator.ValidationErrors) {
        // Use json tag as key
        parts := strings.SplitN(err.Namespace(), ".", 2)
        deepTag := parts[0]
        if len(parts) > 1 {
            deepTag = parts[1]
        }

        errValue := err.Tag()
        if err.Param() != "" {
            errValue += fmt.Sprintf(" - params: (%s)", err.Param())
        }

        errors[deepTag] = errValue
    }
    return errors
}

type Errors map[string]string

func (e Errors) Error() string {
    msg, err := json.Marshal(e)
    if err != nil {
        return err.Error()
    }

    return string(msg)
}
