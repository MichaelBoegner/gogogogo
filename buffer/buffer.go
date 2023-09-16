package buffer

import (
	"bytes"
	"fmt"
)

func Buffering() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Size: %d MB.", 85)
	s := buf.String()
	fmt.Println(s)
	return s
}
