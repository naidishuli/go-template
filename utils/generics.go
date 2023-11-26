package utils

import (
	"math/rand"
	"reflect"
	"time"
)

func PtrTo[value any](v value) *value {
	return &v
}

func ValueOf[T any](value *T) T {
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

func Uniq[T comparable](values []T) []T {
	if values == nil {
		return nil
	}

	l := len(values)
	u := make(map[T]bool, l)
	r := make([]T, 0, l)
	for _, val := range values {
		if _, ok := u[val]; !ok {
			u[val] = true
			r = append(r, val)
		}
	}

	return r
}

func MapKeys[K comparable, V any](values map[K]V) []K {
	if values == nil {
		return nil
	}

	keys := make([]K, 0, len(values))
	for key, _ := range values {
		keys = append(keys, key)
	}

	return keys
}

func RandomPick[T any](values []T) T {
	valuesLen := len(values)
	if valuesLen == 0 {
		var tmp T
		return tmp
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(valuesLen)
	return values[index]
}
