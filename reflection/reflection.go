package reflection

import (
	"reflect"
)

func Walk(x interface{}, fn func(string)) {
	values := getValue(x)
	if values.Kind() == reflect.Slice {
		for i := 0; i < values.Len(); i++ {
			Walk(values.Index(i).Interface(), fn)
		}
	} else {
		for i := 0; i < values.NumField(); i++ {
			field := values.Field(i)

			switch field.Kind() {
			case reflect.String:
				fn(field.String())

			case reflect.Struct:
				Walk(field.Interface(), fn)
			}
		}
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
