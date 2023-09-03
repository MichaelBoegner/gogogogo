package iteration

import "fmt"

func Repeat(letter string) (repeated string) {
	for i := 0; i < 5; i++ {
		if i < 4 {
			repeated += letter + ","
		} else {
			repeated += letter
		}

	}
	fmt.Println(repeated)
	return repeated
}
