package strutures

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("square perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		result := Perimeter(rectangle)
		expected := 40.0

		if (result != expected) {
			t.Errorf("got %.2f want %.2f", result, expected)
		}
	})
}

func TestArea(t *testing.T) {
	// checkArea := func(t *testing.T, shape Shape, expected float64) {
	// 	t.Helper()
	// 	result := shape.Area()
	// 	if result != expected {
	// 		t.Errorf("got %g want %g", result, expected)
	// 	}
	// }
	// t.Run("square area", func(t *testing.T) {
	// 	rectangle := Rectangle{10.0, 10.0}
	// 	expected := 100.0
	// 	checkArea(t, rectangle, expected)
	// })

	// t.Run("circle area", func(t *testing.T) {
	// 	circle := Circle{10.0}
	// 	expected := 314.1592653589793
	// 	checkArea(t, circle, expected)
	// })

	areas := []struct {
		name string
		shape Shape
		expected float64
	}{
		{name: "rectangle", shape: Rectangle{10.0, 10.0}, expected: 100.0},
		{name: "circle", shape: Circle{10.0}, expected: 314.1592653589793},
		{name: "triangle", shape: Triangle{12, 6}, expected: 36.0},
	}

	for _, area := range areas {
		t.Run(area.name, func(t *testing.T) {
			result := area.shape.Area()
			if result != area.expected {
				t.Errorf("%#v got %g want %g", area.name, result, area.expected)
			}
		})
	}
}