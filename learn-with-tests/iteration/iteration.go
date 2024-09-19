package iteration

func Repeat(letter string, times int) string {
	var result string

	for i := 0; i < times; i++ {
		result += letter
	}

	return result
}