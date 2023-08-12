package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Michael")
	fmt.Println(message)
}

