package slices

import "fmt"

func SumSlice(sliced []int) (sum int) {
	for _, num := range sliced {
		sum += num
	}
	return sum
}

func SumAll(slicesToSum ...[]int) (sliceSum []int) {
	numberOfSlices := len(slicesToSum)
	sliceSum = make([]int, numberOfSlices)

	for i, slice := range slicesToSum {
		sliceSum[i] = SumSlice(slice)
	}

	fmt.Println(sliceSum)
	return sliceSum
}
