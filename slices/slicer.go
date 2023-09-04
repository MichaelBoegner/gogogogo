package slices

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

	return sliceSum
}

func SumAllTails(slicesToSum ...[]int) (tailsSum int) {
	for _, slice := range slicesToSum {
		tail := slice[1:]
		tailsSum += SumSlice(tail)
	}

	return tailsSum
}
