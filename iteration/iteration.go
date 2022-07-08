package iteration

func Repeat(character string, limit int) string {
	var repeated string

	for i := 0; i < limit; i++ {
		repeated += character
	}

	return repeated
}
