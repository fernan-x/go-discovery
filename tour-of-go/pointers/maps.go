package pointers

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
	"Yahoo":     {34.00071, -118.25216},
}

func ExampleMap() {
	emptyVertex := Vertex{}
	fmt.Println(emptyVertex)
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(str string) map[string]int {
	words := strings.Fields(str)
	counts := make(map[string]int)

	for i, word:= range words {
		counts[word] = i
	}

	return counts;
}

func ExerciseWordCount() {
	count := WordCount("Hello, World!")
	fmt.Println(count)
}