package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return fmt.Sprintf("%v%v!", spanishHelloPrefix, name)
	}

	if language == french {
		return fmt.Sprintf("%v%v!", frenchHelloPrefix, name)
	}

	return fmt.Sprintf("%v%v!", englishHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
