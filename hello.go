package main

import "fmt"

var standardHumanGreeting, language = "Hello ", "Spanish"

func Hello(standardHumanGreeting, name, language string) string {
	if language == "Spanish" {
		standardHumanGreeting = "Hola "
	}
	if name == "" {
		return standardHumanGreeting + "World!!!"
	}
	return standardHumanGreeting + name + "!!!"
}

func main() {
	fmt.Println(Hello(standardHumanGreeting, "Mike", language))
}
