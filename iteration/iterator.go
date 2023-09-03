package iteration

func Repeat(letter string, times int) (repeated string) {
	if times == 0 {
		times = 5
	}
	for i := 0; i < times; i++ {
		if i < times-1 {
			repeated += letter + ","
		} else {
			repeated += letter
		}
	}

	return repeated
}
