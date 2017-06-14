package main 

import (
	"fmt"
)

type Day int


const (
	MO Day = iota
	WE
	TH
	FR
	SA
	SU
)

var dayName = []string{"Mon", "Tue", "Wed", "Thur", "Fri", "Sat", "Sun"}

func (d Day) String() string {
	return dayName[d]
}

func main() {
	var th Day = SU
	fmt.Println(th)
}