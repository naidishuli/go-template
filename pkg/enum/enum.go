package enum

import "fmt"

type StringType[T comparable] struct {
	def          T
	values       []T
	valuesString []string
}

func NewStringType[T comparable](def T, args ...T) StringType[T] {
	st := StringType[T]{
		def: def,
	}
	for _, arg := range args {
		st.values = append(st.values, arg)
		st.valuesString = append(st.valuesString, fmt.Sprintf("%v", arg))
	}

	return st
}

func (s StringType[T]) Values() []T {
	return s.values
}

func (s StringType[T]) StringValues() []string {
	return s.valuesString
}

func (s StringType[T]) Valid(args T) bool {
	for _, v := range s.values {
		if v == args {
			return true
		}
	}

	return false
}

func (s StringType[T]) Get(arg any) T {
	for _, value := range s.values {
		if fmt.Sprintf("%v", value) == arg {
			return value
		}
	}

	return s.def
}
