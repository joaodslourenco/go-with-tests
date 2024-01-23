package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return fmt.Sprintf("%v%v!", spanishHelloPrefix, name)
	}

	return fmt.Sprintf("%v%v!", englishHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
