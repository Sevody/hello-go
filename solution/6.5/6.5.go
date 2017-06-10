package main

func main() {
	recursive(10)
}
func recursive(n int) {
	if n < 1 {
		return
	}
	println(n)
	recursive(n - 1)
}
