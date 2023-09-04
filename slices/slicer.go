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

func SumAllTails(slicesToSum ...[]int) (tailsSum int) {
	for _, slice := range slicesToSum {
		for j, num := range slice {
			if j > 0 {
				tailsSum += num
			}
		}
	}
	fmt.Println(tailsSum)
	return tailsSum
}
