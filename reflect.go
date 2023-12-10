package utils

import (
	"reflect"
)

var canPointerType = []reflect.Kind{
	reflect.Chan,
	reflect.Func,
	reflect.Map,
	reflect.Pointer,
	reflect.Slice,
	reflect.UnsafePointer,
}

func CanReflectPointer(val reflect.Value) bool {
	kind := val.Kind()

	for _, t := range canPointerType {
		if kind == t {
			return true
		}
	}

	return false
}
