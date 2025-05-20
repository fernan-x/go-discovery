package pointers

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy);

	for y := range dy {
		row := make([]uint8, dx)
		for x := range dx {
			row[x] = uint8((x + y) / 2)
		}
		img[y] = row
	}

	return img
}

func ExercisePointers() {
	Pic(5, 10)
}