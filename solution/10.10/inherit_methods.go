package main

import (
	"fmt"
)

type Base struct {
	id int
}

func (b *Base) Id() int {
	return b.id
}

func (b *Base) SetId(id int) {
	b.id = id
}

type Person struct {
	Base 
	FirstName string
	LastName string
}

type Employee struct {
	Person
	salary int
}

func main() {
	baka := &Employee{Person{Base{67}, "lu", "kino"}, 5}
	fmt.Println(baka.Id())
	baka.SetId(9)
	fmt.Println(baka.Id())
	
}