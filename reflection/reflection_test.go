package reflection

import "testing"

func TestWalk(t *testing.T) {
	t.Run("walk(x interface{}, fn func(string)) which takes a struct x and calls fn for all strings fields found inside.", func(t *testing.T) {
		x := struct {
			name   string
			place  string
			place2 string
			age    int
		}{name: "frank"}

		var got []string

		fn := func(input string) {
			got = append(got, input)
		}

		Walk(x, fn)

		if len(got) != 1 {
			t.Errorf("there should only be %d calls, but got %d number of calls", 1, len(got))
		}
	})
}
