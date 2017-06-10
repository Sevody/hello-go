package main

func main() {
	for i := 20; i > 0; i-- {
		if i > 10 {
			if i == 15 {
				continue
			}
		}
		println(i)
	}
}
