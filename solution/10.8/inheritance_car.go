package main

import (
	"fmt"
)

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	wheelCount int
	Engine
}

func (c *Car) numberOfWheels() int {
	return c.wheelCount
}

func (c *Car) Start() {
	fmt.Println("start")
}

func (c *Car) Stop() {
	fmt.Println("stop")
}

func (c *Car) GoTOWorkIn() {
	c.Start()
	c.Stop()
}
type Mercedes struct {
	Car
}

func (m *Mercedes) sayHiToMerkel() {
 fmt.Println("hi")
}

func main() {
	m := &Mercedes{Car{4, nil}}
	m.GoTOWorkIn()
}