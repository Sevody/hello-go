package main

func fib() func(int) int {
	a, b := 1, 1
	return func(n int) int {
		a, b = b, a+b
		return b
	}
}
func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		println(f(i + 2))
	}
}
