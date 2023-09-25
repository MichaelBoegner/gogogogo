package _select

import (
	"fmt"
	"net/http"
	"time"
)

var configuredTimeOut = time.Duration(10 * time.Millisecond)

func Racer(URLA, URLB string) (string, error) {
	return ConfigurationRacer(URLA, URLB, configuredTimeOut)
}

func ConfigurationRacer(URLA, URLB string, timeout time.Duration) (string, error) {
	select {
	case <-ping(URLA):
		return URLA, nil
	case <-ping(URLB):
		return URLB, nil
	case <-time.After(timeout):
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
