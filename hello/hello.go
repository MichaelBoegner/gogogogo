package hello

import "fmt"

const language = "French"

func Hello(name, language string) string {
	standardHumanGreeting := ""

	switch language {
	case "Spanish":
		standardHumanGreeting = "Hola "
	case "French":
		standardHumanGreeting = "Bonjour "
	case "English":
		standardHumanGreeting = "Hello "
	}

	if name == "" {
		return standardHumanGreeting + "World!!!"
	}
	return standardHumanGreeting + name + "!!!"
}

func RunHello() {
	fmt.Println(Hello("Mike", language))
}
