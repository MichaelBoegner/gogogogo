package _select

import "testing"

func Test_Select(t *testing.T) {
	t.Run("take two URLs and race them against each other using GET. First to return site wins.", func(t *testing.T) {
		URLFast := "https://cmboegner.com"
		URLSlow := "https://google.com"

		got := Racer(URLFast, URLSlow)
		want := URLFast

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
