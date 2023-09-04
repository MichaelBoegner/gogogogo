package slices

import "fmt"

func Slice(sliced []int) (sum int) {
	for _, num := range sliced {
		sum += num
	}
	return sum
}

func SumAll(sliceA []int, sliceB []int) (sliceSum []int) {
	sliceSumA, sliceSumB := 0, 0

	for _, num := range sliceA {
		sliceSumA += num
	}
	sliceSum = append(sliceSum, sliceSumA)
	for _, num := range sliceB {
		sliceSumB += num
	}
	sliceSum = append(sliceSum, sliceSumB)
	fmt.Println(sliceSum)
	return sliceSum
}
