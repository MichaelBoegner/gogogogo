package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("Search method is able to return value by passing an existing key to dictionary", func(t *testing.T) {
		dictionary := Dictionary{"firstKey": "firstValue"}

		got := dictionary.Search("firstKey")
		want := "firstValue"

		assertSearch(t, got, want)
	})
	t.Run("Search method is able to return an error when passing non-existing key to dictionary", func(t *testing.T) {
		dictionary := Dictionary{"firstKey": "firstValue"}

		got := dictionary.Search("farts")
		want := "Key does not exist in Dictionary"

		assertSearch(t, got, want)
	})

}

func assertSearch(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}
