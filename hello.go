package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const norwegian = "Norwegian"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const norwegianHelloPrefix = "Hei, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case norwegian:
		prefix = norwegianHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Ulf", "Norwegian"))
}
