package structutil

import "reflect"

func MergeStructs[T any](defaultStruct, overrideStruct T) T {
	defaultValue := reflect.ValueOf(defaultStruct)
	overrideValue := reflect.ValueOf(overrideStruct)
	result := reflect.New(defaultValue.Type()).Elem()

	for i := 0; i < defaultValue.NumField(); i++ {
		field := overrideValue.Field(i)
		if !field.IsZero() {
			result.Field(i).Set(field)
		} else {
			result.Field(i).Set(defaultValue.Field(i))
		}
	}

	return result.Interface().(T)
}
