package countdown

import (
	"bytes"
	"fmt"
	"time"
)

var buffer = &bytes.Buffer{}

func Countdown(out *bytes.Buffer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintf(out, "Go!")
}
