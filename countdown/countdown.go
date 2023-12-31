package countdown

import (
	"fmt"
	"io"
	"os"
	"time"
)

var buffer = os.Stdout

func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintf(out, "Go!")
}
