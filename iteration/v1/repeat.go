package iteration

const repeatCount int = 5

// Repeat ...
func Repeat(character string) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += "a"
	}
	return repeated
}
