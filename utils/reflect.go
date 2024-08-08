package utils

import "reflect"

func TypeName(typ interface{}) string {
    if typ == nil {
        return ""
    }

    switch v := typ.(type) {
    case string:
        return v
    case *string:
        return ValueOf(v)
    }

    v := reflect.ValueOf(typ)

    if v.Kind() == reflect.Ptr {
        return reflect.Indirect(v).Type().Name()
    }

    a := v.Type().Name()
    return a
}

func IsConvertibleTo(a, b any) bool {
    typeOfA := reflect.TypeOf(a)
    typeOfB := reflect.TypeOf(b)

    return typeOfA.ConvertibleTo(typeOfB)
}
