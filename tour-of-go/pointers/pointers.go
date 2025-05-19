package pointers

import "fmt"

func ExamplePointer() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

func ExampleStruct() {
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)

	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func ExampleStructLiteral() {
	type Vertex struct {
		X, Y int
	}

	v1 := Vertex{1, 2}
	v2 := Vertex{X: 1}
	v3 := Vertex{}
	p := &Vertex{1, 4}

	fmt.Println(v1, v2, v3, p.X)
}

func ExampleArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func ExampleSlice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	// Literal
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	q[5] = 100
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func ExampleSliceLength() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func ExampleSliceNil() {
	var s []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(s), cap(s), s)
	if s == nil {
		fmt.Println("nil!")
	}
}

func ExampleMakeSlice() {
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)
}

func ExampleSliceOfSlices() {
	board := [][]string{
		[]string{"_", "_", "_" },
		[]string{"_", "_", "_" },
		[]string{"_", "_", "_" },
	}

	// fmt.Printf("%v\n", board)
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for _, row := range board {
		fmt.Println(row)
	}
}