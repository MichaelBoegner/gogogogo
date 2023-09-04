package slices

import "fmt"

func SumSlice(sliced []int) (sum int) {
	for _, num := range sliced {
		sum += num
	}
	return sum
}

func SumAll(slicesToSum ...[]int) (sliceSum []int) {
	for _, slice := range slicesToSum {
		sliceSum = append(sliceSum, SumSlice(slice))
	}

	fmt.Println(sliceSum)
	return sliceSum
}
