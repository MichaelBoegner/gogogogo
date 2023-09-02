package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello(standardHumanGreeting, "Mike", "English")
		want := "Hello Mike!!!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is input", func(t *testing.T) {
		got := Hello(standardHumanGreeting, "", "English")
		want := "Hello World!!!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say it in Spanish", func(t *testing.T) {
		got := Hello(standardHumanGreeting, "Mike", "Spanish")
		want := "Hola Mike!!!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, but we want %q", got, want)
	}
}
