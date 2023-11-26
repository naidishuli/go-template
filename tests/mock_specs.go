package tests

import (
	"go.uber.org/mock/gomock"
	"reflect"
)

type MockSpecs struct {
	Calls        int
	Args         []interface{}
	SetArgs      map[int]interface{}
	ReturnValues []interface{}
	DoAndReturn  interface{}
	Do           interface{}
}

type MockCase map[string][]MockSpecs

func SetDepExpectation(dep any, mockCase MockCase) {
	for fn, specs := range mockCase {
		for _, spec := range specs {
			if spec.Calls < 1 {
				continue
			}

			call := gomockCallFromDep(dep, fn, spec.Args...)
			if call == nil {
				continue
			}

			call = call.Times(spec.Calls)

			if len(spec.SetArgs) > 0 {
				for index, setArg := range spec.SetArgs {
					call = call.SetArg(index, setArg)
				}
			}

			if len(spec.ReturnValues) > 0 {
				call = call.Return(spec.ReturnValues...)
			}

			if spec.Do != nil {
				call = call.Do(spec.DoAndReturn)
			}

			if spec.DoAndReturn != nil {
				call = call.DoAndReturn(spec.DoAndReturn)
			}
		}
	}
}

func gomockCallFromDep(dep any, fn string, args ...any) *gomock.Call {
	val := reflect.ValueOf(dep)
	method := val.MethodByName("EXPECT")
	if !method.IsValid() {
		return nil
	}

	rVal := method.Call([]reflect.Value{})
	mockStrc := rVal[0].Interface()

	val = reflect.ValueOf(mockStrc)
	method = val.MethodByName(fn)
	if !method.IsValid() {
		return nil
	}

	var values []reflect.Value
	for _, arg := range args {
		values = append(values, reflect.ValueOf(arg))
	}

	return method.Call(values)[0].Interface().(*gomock.Call)
}
