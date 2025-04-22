package main

import "packages/flowcontrol"

// func chapter1() {
// 	// Importing packages
// 	fmt.Println("Hello, World!", rand.Intn(10))
// 	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
// 	fmt.Println(math.Pi)
// 	fmt.Printf("The sum of 42 and 13 is %d.\n", utils.Add(42, 13))

// 	// Function values
// 	a, b := utils.Swap("Hello", "World")
// 	fmt.Println(a, b)

// 	// Variables with initializers
// 	c, python, java := true, false, "no!"
// 	var i int
// 	fmt.Println(i, c, python, java)

// 	// Type conversions
// 	i = 42
// 	var f float64 = float64(i)
// 	u := uint(f)
// 	fmt.Println(i, f, u)

// 	// Constants
// 	const Pi = 3.14
// 	fmt.Println("Happy", Pi, "Day")
// }

// func playingWithPackages() {
// 	utils.Example()
// 	res := utils.Add(1, 2)
// 	fmt.Println(res)
// 	a, b := utils.Swap("Hello", "World")
// 	fmt.Println(a, b)
// }

func playingWithFlowControl() {
	// flowcontrol.ExampleFor()
	// flowcontrol.ExampleIf()
	flowcontrol.ExerciseLoop()
}

func main() {
	// playingWithPackages()
	playingWithFlowControl()
}
