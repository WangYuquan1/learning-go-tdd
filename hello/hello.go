package main

import "fmt"

const spanish = "Spanish"
const engilishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = engilishHelloPrefix
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
