package main

import (
	"fmt"
)

func main() {
	t := []float32{1.2,3.2,5.4}
	fmt.Println(Sum(t))
}

func Sum(arrF []float32) (s float32) {
	for _, v := range arrF {
		s += v
	}
	return
} 

