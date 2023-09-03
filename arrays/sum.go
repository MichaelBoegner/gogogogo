package arrays

func Sum(intArray []int) (sum int) {
	for _, num := range intArray {
		sum += num
	}
	return sum
}
