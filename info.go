package utils

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func QueryName(f any) string {
	val := reflect.ValueOf(f)
	if !CanReflectPointer(val) {
		return fmt.Sprint(val)
	}

	return runtime.FuncForPC(val.Pointer()).Name()
}

func NameOfPackage() string {
	tmpName := QueryName(NameOfPackage)

	return tmpName[:strings.LastIndex(tmpName, ".")]
}
