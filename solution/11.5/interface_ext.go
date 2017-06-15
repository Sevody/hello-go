package main

import "fmt"

type Shaper interface {
	Area() float32
}

type AreaI interface {
	Area() float32
}

type Triangle struct {
	width, height float32
}

func (t Triangle) Area() float32 {
	return t.height * t.width * 0.5
}

type Square struct {
	side float32
	test int
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
	test int
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * 3.14
}
func main() {
	r := Rectangle{5, 3, 9}
	q := &Square{5, 9}
	c := Circle{3}
	shapes := []Shaper{r, q, c}
	fmt.Println("Looping through shapes for area")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
