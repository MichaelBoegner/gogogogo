package _select

import (
	"net/http"
	"time"
)

func Racer(URLA, URLB string) string {
	startA := time.Now()
	_, _ = http.Get(URLA)
	respADuration := time.Since(startA)

	startB := time.Now()
	_, _ = http.Get(URLB)
	respBDuration := time.Since(startB)

	if respADuration < respBDuration {
		return URLA
	} else {
		return URLB
	}
}
