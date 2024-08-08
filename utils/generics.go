package utils

import (
    "fmt"
    "reflect"
    "strings"
)

type BasicValue interface {
    int | string
}

func PtrTo[value any](v value) *value {
    return &v
}

func ValueOf[T comparable](value *T) T {
    if value == nil {
        return reflect.Zero(getType(value)).Interface().(T)
    }

    return *value
}

func getType(a any) reflect.Type {
    v := reflect.ValueOf(a)
    if v.Kind() == reflect.Ptr {
        return v.Type().Elem()
    }

    return v.Type()
}

func SliceOfAny[T any](values []T) []any {
    a := make([]any, 0, len(values))
    for _, v := range values {
        a = append(a, v)
    }

    return a
}

func Join[T any](elems []T, sep string) string {
    n := len(elems)

    switch n {
    case 0:
        return ""
    case 1:
        return fmt.Sprintf("%v", elems[0])
    }

    var b strings.Builder
    b.Grow(n)
    b.WriteString(fmt.Sprintf("%v", elems[0]))
    for _, s := range elems[1:] {
        b.WriteString(sep)
        b.WriteString(fmt.Sprintf("%v", s))
    }
    return b.String()
}

func Ternary[T any](condition bool, a, b T) T {
    if condition {
        return a
    }

    return b
}

func ArrayIncludes[T comparable](arr []T, value T) bool {
    for _, v := range arr {
        if v == value {
            return true
        }
    }

    return false
}
