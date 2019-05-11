package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const kannada = "Kannada"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const kannadaHelloPrefix = "Namaskara, "

// Hello returns a personalised greeting in a given language
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case kannada:
		prefix = kannadaHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Print(Hello("world", ""))
}
