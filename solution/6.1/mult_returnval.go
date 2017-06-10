package function

func Returnval(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}
func Returnval2(a, b int) (sum, mul, diff int) {
	sum = a + b
	mul = a * b
	diff = a - b
	return
}
