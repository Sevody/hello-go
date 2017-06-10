package main

import "fmt"

func main() {
	month := 6
	s := Season(month)
	fmt.Printf("this month is %s", s)
}

func Season(m int) (s string) {
	switch m {
	case 3, 4, 5:
		s = "Spring"
	case 6, 7, 8:
		s = "Summer"
	case 9, 10, 11:
		s = "Fall"
	case 12, 1, 2:
		s = "Winter"
	}
	return s
}
