package main

import "fmt"

var standardHumanGreeting, language = "Hello ", "English"

func Hello(standardHumanGreeting, name, language string) string {
	if language == "Spanish" {
		standardHumanGreeting = "Hola "
	}

	if language == "French" {
		standardHumanGreeting = "Bonjour "
	}

	if name == "" {
		return standardHumanGreeting + "World!!!"
	}
	return standardHumanGreeting + name + "!!!"
}

func main() {
	fmt.Println(Hello(standardHumanGreeting, "Mike", language))
}
