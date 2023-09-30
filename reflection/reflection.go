package reflection

import (
	"reflect"
)

func Walk(x interface{}, fn func(string)) {
	values := reflect.ValueOf(x)
	for i := 0; i < values.NumField(); i++ {
		field := values.Field(i)
		fn(field.String())
	}
}
