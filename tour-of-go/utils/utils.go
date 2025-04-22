package utils

import "fmt"

func Example() {
	fmt.Println("When's Saturday?")
}

func Add(x, y int) int {
	return x + y
}

func Swap(x, y string) (string, string) {
	return y, x
}
