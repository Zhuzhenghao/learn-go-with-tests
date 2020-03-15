package iteration

// Repeat ...
func Repeat(character string, count int) (repeated string) {
	for i := 0; i < count; i++ {
		repeated += "a"
	}
	return repeated
}
