package main

import "fmt"

const vatRate = 20.0

func main() {
	printVatRate((vatRate))
}

func printVatRate(vatRate float32) {
	fmt.Printf("%.2f%%\n", vatRate)
}
