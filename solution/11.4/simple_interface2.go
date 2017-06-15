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

type RSimple struct {
	value int
}

func (s *RSimple) Set(arg int) {
	s.value = arg
}

func (s *RSimple) Get() int {
	return s.value
}

func fi(is ...interface{}) {
	for _, s := range is {
		switch s.(type) {
			case *Simple:
				fmt.Println("Simple struct:", s)
			case *RSimple:
				fmt.Println("RSimple struct:", s)
		}
	}
}

func main() {
	s := &Simple{3}
	var ser Simpler
	ser = s
	ser.Set(5)
	fmt.Println(ser.Get())
	s2 := &RSimple{4}
	fi(s, s2)
}