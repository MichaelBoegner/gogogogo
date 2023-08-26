package main

import (
	"fmt"
	"sync"
)

func generateNumbers(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	numbersSlice := make([]int, 5)
	for i := 1; i < len(numbersSlice); i++ {
		fmt.Printf("sending %d to channel\n", i)
		ch <- i
		fmt.Println("i:", i)
		fmt.Println("ch:", ch)
	}
}

func secondGoroutine(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("reading %d from channel\n", num)
	}

}

func main() {
	var wg sync.WaitGroup
	//fmt.Println("wg:", wg)
	fmt.Println("&wg:", &wg)

	numberChan := make(chan int)

	wg.Add(2)
	go generateNumbers(numberChan, &wg)
	secondGoroutine(numberChan, &wg)

	close(numberChan)

	fmt.Println("Waiting for goroutines to finish . . .")
	wg.Wait()
	fmt.Println("Done!")
}
