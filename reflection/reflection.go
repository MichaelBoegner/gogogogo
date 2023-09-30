package reflection

import (
	"reflect"
)

func Walk(x interface{}, fn func(string)) {
	reflected := reflect.ValueOf(x)
	field := reflected.Field(0)
	fn(field.String())
}
