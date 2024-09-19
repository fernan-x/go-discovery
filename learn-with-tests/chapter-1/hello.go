package main

import "fmt"

const (
	english = "English"
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "¡Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, lang string) string {
	if (name == "") {
		name = "World"
	}
	name = name + "!"

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world
func main() {
	fmt.Println(Hello("world", ""))
}