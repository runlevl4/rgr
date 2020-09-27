package iteration

// Repeat takes an input string and repeats it 5 times.
func Repeat(x string, count int) string {
	var tmp string

	for i := 0; i < count; i++ {
		tmp += x
	}
	return tmp
}