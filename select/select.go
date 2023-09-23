package _select

import (
	"net/http"
	"time"
)

func Racer(URLA, URLB string) string {
	aTime := responseTime(URLA)
	bTime := responseTime(URLB)

	if aTime < bTime {
		return URLA
	} else {
		return URLB
	}
}

func responseTime(URL string) time.Duration {
	start := time.Now()
	_, _ = http.Get(URL)
	return time.Since(start)
}
