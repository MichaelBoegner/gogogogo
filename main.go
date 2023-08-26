package main

import (
	"fmt"
)

func ranges() {
	var elems = []int{1, 2, 3, 4, 5}
	sum := 0

	for i, elem := range elems {
		sum += elem
		fmt.Println("i", i)
		fmt.Println("elem", elem)
	}
	fmt.Println("sum", sum)
}

func main() {
	ranges()
}
