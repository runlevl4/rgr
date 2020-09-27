package iteration

// Repeat takes an input string and repeats it 5 times.
func Repeat(x string) string {
	var tmp string

	for i := 0; i < 5; i++ {
		tmp += x
	}
	return tmp
}