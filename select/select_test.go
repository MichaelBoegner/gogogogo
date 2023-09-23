package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("take two URLs and race them against each other using GET. First to return site wins.", func(t *testing.T) {
		fastServer := serverCreator(0 * time.Millisecond)
		slowServer := serverCreator(20 * time.Millisecond)

		defer fastServer.Close()
		defer slowServer.Close()

		URLFast := fastServer.URL
		URLSlow := slowServer.URL

		got := Racer(URLFast, URLSlow)
		want := URLFast

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("returns an error if server takes more than 10 seconds to return response", func(t *testing.T) {
		fastServer := serverCreator(0 * time.Second)
		slowServer := serverCreator(11 * time.Second)

		defer fastServer.Close()
		defer slowServer.Close()

		URLFast := fastServer.URL
		URLTimeOut := slowServer.URL

		_, err := Racer(URLFast, URLTimeOut)

		if err == nil {
			t.Errorf("expected an error but got nil")
		}
	})
}

func serverCreator(timeDelay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(timeDelay)
		w.WriteHeader(http.StatusOK)
	}))
}
