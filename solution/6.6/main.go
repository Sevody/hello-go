package main

func main() {
	ret := recursive(30)
	println(ret)
}

func recursive(n float64) float64 {
	if n <= 1 {
		return 1
	}
	return n * recursive(n-1)
}
