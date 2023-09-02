package main

import "fmt"

const standardHumanGreeting = "Hello "

func Hello(standardHumanGreeting, name string) string {
	if name == "" {
		return standardHumanGreeting + "World!!!"
	}
	return standardHumanGreeting + name + "!!!"
}

func main() {
	fmt.Println(Hello(standardHumanGreeting, ""))
}
