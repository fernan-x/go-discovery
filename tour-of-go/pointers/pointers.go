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