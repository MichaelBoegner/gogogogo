package dictionary

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {

	t.Run("Search method is able to return value by passing an existing key to dictionary", func(t *testing.T) {
		dictionary := Dictionary{"firstKey": "firstValue"}

		got, _ := dictionary.Search("firstKey")
		want := "firstValue"

		assertSearch(t, got, want)
	})
	t.Run("Search method is able to return an error when passing non-existing key to dictionary", func(t *testing.T) {
		dictionary := Dictionary{"firstKey": "firstValue"}

		_, got := dictionary.Search("farts")

		assertSearchError(t, got, ErrNoKey)
	})
}

func TestSubmit(t *testing.T) {
	t.Run("Submit method should submit a new entry into the dictionary", func(t *testing.T) {
		dictionary := Dictionary{"firstKey": "firstValue"}
		dictionary.Submit("secondKey", "secondValue")

		got, err := dictionary.Search("secondKey")
		want := "secondValue"

		assertSubmit(t, got, want, err)
	})
	t.Run("Submit method should not overwrite existing key:value pairs", func(t *testing.T) {
		existingKey := "existingKey"
		dictionary := Dictionary{existingKey: "existingValue"}

		got := dictionary.Submit(existingKey, "newValue")
		want := errors.New("cannot use existing key to overwrite value")

		if !errors.Is(got, want) {
			t.Errorf("got %q, but wanted %q", got, want)
		}
	})
}

func assertSearch(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}

func assertSearchError(t testing.TB, got, want error) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}

func assertSubmit(t testing.TB, got, want string, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("Fatal test flaw. Should not have returned the error:", err)
	}

	assertSearch(t, got, want)
}
