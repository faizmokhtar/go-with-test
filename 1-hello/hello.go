package main

import "fmt"

const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	const englishHelloPrefix = "Hello, "
	const spanishHelloPrefix = "Hola, "
	const franceHelloPrefix = "Bonjour, "

	switch language {
	case french:
		prefix = franceHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
