package iteration

import "fmt"

func Repeat(letter string, times int) (repeated string) {
	if times == 0 {
		for i := 0; i < 5; i++ {
			if i < 4 {
				repeated += letter + ","
			} else {
				repeated += letter
			}
		}
	} else {
		for i := 0; i < times; i++ {
			if i < times-1 {
				repeated += letter + ","
			} else {
				repeated += letter
			}
		}
	}
	fmt.Println(repeated)
	return repeated
}
