package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{"frank"},
			ExpectedCalls: []string{"frank"},
		},
	}

	for _, test := range cases {
		t.Run("walk(x interface{}, fn func(string)) which takes a struct x and calls fn for all strings fields found inside.", func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})
			values := reflect.ValueOf(test.Input)
			length := values.NumField()
			if len(got) != length {
				t.Errorf("there should only be %d calls, but got %d number of calls", length, len(got))
			}
		})
	}

}
