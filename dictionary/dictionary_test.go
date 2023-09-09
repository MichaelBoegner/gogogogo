package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("Search() is able to return value by passing key to map", func(t *testing.T) {
		dictionary := map[string]string{"firstKey": "firstValue"}

		got := Search(dictionary, "firstKey")
		want := "firstValue"

		if got != want {
			t.Errorf("got %q, but wanted %q", got, want)
		}
	})

}
