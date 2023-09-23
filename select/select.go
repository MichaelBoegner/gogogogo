package _select

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(URLA, URLB string) (string, error) {
	select {
	case <-ping(URLA):
		return URLA, nil
	case <-ping(URLB):
		return URLB, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", URLA, URLB)
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
