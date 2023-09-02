package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello(standardHumanGreeting, "Mike")
		want := "Hello Mike!!!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is input", func(t *testing.T) {
		got := Hello(standardHumanGreeting, "")
		want := "Hello World!!!"

		if got != want {
			t.Errorf("got %q, but we want %q", got, want)
		}
	})

}
