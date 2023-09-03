package slices

func Slice(sliced []int) (sum int) {
	for _, num := range sliced {
		sum += num
	}
	return sum
}
