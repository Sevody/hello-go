package main

import (
	"fmt"
)

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	value int
}

func (s *Simple) Set(arg int) {
	s.value = arg
}

func (s *Simple) Get() int {
	return s.value
}

func main() {
	s := &Simple{3}
	var ser Simpler
	ser = s
	ser.Set(5)
	fmt.Println(ser.Get())
}