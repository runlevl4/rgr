package iteration

func Repeat(x string) string {
	var tmp string

	for i := 0; i < 5; i++ {
		tmp = tmp + x
	}
	return tmp
}