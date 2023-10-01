package reflection

import (
	"fmt"
	"reflect"
)

func Walk(x interface{}, fn func(string)) {
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
		fmt.Println("this is reflect.string", val)
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice:
		numberOfValues = val.Len()
		getField = val.Index
	}

	for i := 0; i < numberOfValues; i++ {
		Walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	values := reflect.ValueOf(x)
	switch values.Kind() {
	case reflect.Pointer:
		values = values.Elem()
	}

	return values
}
