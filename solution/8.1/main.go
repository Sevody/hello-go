package main

import (
	"fmt"
)

func main() {
	week := make(map[int]string)
	week[1] = "Monday"
	week[2] = "Tuesday"
	week[3] = "Wensday"
	week[4] = "Thurday"
	week[5] = "Friday"
	week[6] = "Saturday"
	week[7] = "Sunday"

	for _, v := range week {
		fmt.Println("today is", v)
	}
}