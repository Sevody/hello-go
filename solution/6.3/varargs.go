package main

func main() {
	s := []string{"aa", "bb", "cc"}
	Pr(s...)
}

func Pr(s ...string) {
	for _, i := range s {
		println(i)
	}
}
