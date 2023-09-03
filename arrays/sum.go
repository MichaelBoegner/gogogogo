package arrays

func Sum(intArray [3]int) (sum int) {
	for _, num := range intArray {
		sum += num
	}
	return sum
}
