package _select

import (
	"net/http"
)

func Racer(URLA, URLB string) (string, error) {
	select {
	case <-ping(URLA):
		return URLA, nil
	case <-ping(URLB):
		return URLB, nil
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
