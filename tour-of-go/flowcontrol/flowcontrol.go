package flowcontrol

import (
	"fmt"
	"math"
	"math/rand"
)

func ExampleFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)
}

func ExampleIf() {
	if v := rand.Intn(2); v == 0 {
		fmt.Println("true", v)
	} else {
		fmt.Println("false", v)
	}
}

func sqrt(x float64) float64 {
	z := 1.0
	for math.Abs(z*z-x) > 1e-10 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func ExerciseLoop() {
	fmt.Println(sqrt(2))
}