package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name     string
	Pet      string
	Hometown Hometown
	Age      int
}

type Hometown struct {
	Name string
	Lat  float64
	Long float64
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name:          "struct with one string field",
			Input:         Person{Name: "Frank"},
			ExpectedCalls: []string{"Frank"},
		},
		{
			Name:          "struct with two string field",
			Input:         Person{Name: "Frank", Pet: "Franko"},
			ExpectedCalls: []string{"Frank", "Franko"},
		},
		{
			Name:          "struct with one string and one int field",
			Input:         Hometown{Name: "Mexico", Lat: 33.56754},
			ExpectedCalls: []string{"Mexico"},
		},
		{
			Name:          "struct with subsets",
			Input:         Person{Name: "Frank", Pet: "Franko", Hometown: Hometown{Name: "Mexico", Lat: 33.56754}},
			ExpectedCalls: []string{"Frank", "Franko", "Mexico"},
		},
		{
			Name:          "pointer",
			Input:         &Person{Name: "Frank", Pet: "Franko"},
			ExpectedCalls: []string{"Frank", "Franko"},
		},
		{
			Name: "slice",
			Input: []Person{
				{Name: "Frank", Pet: "Franko"},
				{Name: "Matt", Pet: "Matto"},
			},
			ExpectedCalls: []string{"Frank", "Franko", "Matt", "Matto"},
		},
		{
			Name: "maps",
			Input: map[string]string{
				"Birthplace": "Austria",
				"Food":       "Bananas",
			},
			ExpectedCalls: []string{"Austria", "Bananas"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
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
