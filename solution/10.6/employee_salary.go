package main

import (
	"fmt"
)

type employee struct {
	salary float32
}

func (em *employee) giveRaise(raise float32) {
	em.salary = em.salary * raise
}

func main() {
	tom := &employee{1000}
	tom.giveRaise(1.0000002)
	fmt.Println(tom.salary)
}