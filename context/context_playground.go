package main

import (
	"context"
	"fmt"
	"time"
)

func dataSender(ctx context.Context) {
	ctx, cancelCtx := context.WithCancel(ctx)
	numChan := make(chan int)
	go dataHandler(ctx, numChan)
	for num := 1; num <= 5; num++ {
		numChan <- num
	}

	cancelCtx()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("dataSender finished")
}

func dataHandler(ctx context.Context, numChan <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\n ctx.Done processed: %v \n", ctx)
			if err := ctx.Err(); err != nil {
				fmt.Printf("The ctx.Done() processed because: %v\n", err)
			}
			return
		case result := <-numChan:
			fmt.Printf("Printing result: %d \n", result)
		}
	}
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "parentKey", "parentValue")
	dataSender(ctx)
}
