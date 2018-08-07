package iteration

// Repeat takes a string and integers, n and returns a repeated string of n times
func Repeat(character string, frequency int) (repeated string) {
	for i := 0; i < frequency; i++ {
		repeated += character
	}
	return
}
