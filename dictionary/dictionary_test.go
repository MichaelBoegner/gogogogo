package dictionary

import (
	"errors"
	"testing"
)

func TestDictionary(t *testing.T) {
	t.Run("Search method is able to return value by passing an existing key to dictionary", func(t *testing.T) {
		dictionary := Dictionary{"firstKey": "firstValue"}

		got, _ := dictionary.Search("firstKey")
		want := "firstValue"

		assertSearch(t, got, want)
	})
	t.Run("Search method is able to return an error when passing non-existing key to dictionary", func(t *testing.T) {
		dictionary := Dictionary{"firstKey": "firstValue"}

		_, got := dictionary.Search("farts")

		assertError(t, got, ErrNoKey)
	})
	t.Run("Submit method should submit a new entry into the dictionary", func(t *testing.T) {
		dictionary := Dictionary{"firstKey": "firstValue"}
		dictionary.Submit("secondKey", "secondValue")

		got, _ := dictionary.Search("secondKey")
		want := "secondValue"

		assertSearch(t, got, want)
	})
}

func assertSearch(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if !errors.Is(want, got) {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}
