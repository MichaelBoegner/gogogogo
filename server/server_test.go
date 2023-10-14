package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("getPlayers should get player from endpoint /players/{name}", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Mike", nil)
		response := httptest.NewRecorder()

		GetPlayers(response, request)

		got := response.Body.String()
		want := "Mike"

		if got != want {
			t.Errorf("got %q, but wanted %q", got, want)
		}
	})
}
