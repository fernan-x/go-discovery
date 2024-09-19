package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if (name == "") {
		name = "World"
	}
	return englishHelloPrefix + name + "!"
}

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world
func main() {
	fmt.Println(Hello("world"))
}