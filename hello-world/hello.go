package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const portuguese = "Portuguese"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const portugueseHelloPrefix = "Olá, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	var prefix string

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return fmt.Sprintf("%v%v!", prefix, name)
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
