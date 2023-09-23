package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("take two URLs and race them against each other using GET. First to return site wins.", func(t *testing.T) {
		fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))

		URLFast := fastServer.URL
		URLSlow := slowServer.URL

		got := Racer(URLFast, URLSlow)
		want := URLFast

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}

		fastServer.Close()
		slowServer.Close()
	})
}
