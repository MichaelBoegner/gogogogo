package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Mike", "English")
		want := "Hello Mike!!!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is input", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello World!!!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello("Mike", "Spanish")
		want := "Hola Mike!!!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("Mike", "French")
		want := "Bonjour Mike!!!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, but we want %q", got, want)
	}
}
