package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("Search method is able to return value by passing an existing key to dictionary", func(t *testing.T) {
		dictionary := Dictionary{"firstKey": "firstValue"}

		got, err := dictionary.Search("firstKey")
		want := "firstValue"

		assertSearch(t, got, want, err)
	})
	t.Run("Search method is able to return an error when passing non-existing key to dictionary", func(t *testing.T) {
		dictionary := Dictionary{"firstKey": "firstValue"}

		got, err := dictionary.Search("farts")
		want := "Key does not exist in Dictionary"

		assertSearch(t, got, want, err)
	})

}

func assertSearch(t testing.TB, got, want string, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("Fatal fucking error. Stop your shit now.")
	}
	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}
