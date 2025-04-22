package flowcontrol

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
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

func ExampleSwitch() {
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func exampleDeferOtherFunction() {
	defer fmt.Println("defer other function")
	fmt.Println("end other function")
}

func ExampleDefer() {
	defer fmt.Println("defer")
	fmt.Println("end")
	exampleDeferOtherFunction()
}