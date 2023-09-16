package countdown

import (
	"bytes"
	"fmt"
)

var buffer = &bytes.Buffer{}

func Countdown(out *bytes.Buffer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprintf(out, "Go!")
}

func main() {
	Countdown(buffer)
}
