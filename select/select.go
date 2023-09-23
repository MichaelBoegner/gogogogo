package _select

import "net/http"

func Racer(URLOne, URLTwo string) string {
	_, _ = http.Get(URLOne)
	_, _ = http.Get(URLTwo)

	return ""
}
