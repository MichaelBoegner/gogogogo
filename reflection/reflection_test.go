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
		{
			Name: "struct with two string field",
			Input: struct {
				Name  string
				Place string
			}{"Frank", "Mexico"},
			ExpectedCalls: []string{"Frank", "Mexico"},
		},
	}

	for _, test := range cases {
		t.Run("walk(x interface{}, fn func(string)) which takes a struct x and calls fn for all strings fields found inside.", func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, but wanted %v", got, test.ExpectedCalls)
			}
		})
	}

}
