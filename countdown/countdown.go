package countdown

import (
	"bytes"
	"fmt"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type Sleeper interface {
	Sleep()
}

func Countdown(buffer *bytes.Buffer) {
	fmt.Fprintf(buffer, "3")
}
