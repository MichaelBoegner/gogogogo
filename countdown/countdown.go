package countdown

import (
	"bytes"
	"fmt"
)

var buffer = &bytes.Buffer{}

func Countdown(out *bytes.Buffer) {
	fmt.Fprintf(out, `3,
2,
1,
Go!`)
}

func main() {
	Countdown(buffer)
}
