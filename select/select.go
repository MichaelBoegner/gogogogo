package _select

import (
	"net/http"
)

func Racer(URLA, URLB string) string {
	select {
	case <-ping(URLA):
		return URLA
	case <-ping(URLB):
		return URLB
	}
}

func ping(URL string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, _ = http.Get(URL)
		close(ch)
	}()
	return ch
}
