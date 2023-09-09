package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("Search() is able to return value by passing key to map", func(t *testing.T) {
		dictionary := Dictionary{"firstKey": "firstValue"}

		got := Search(dictionary, "firstKey")
		want := "firstValue"

		assertSearch(t, got, want)
	})

}

func assertSearch(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}
