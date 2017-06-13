package main

func main() {
	re := &Rectangle{10, 5}
	println("re Area is", re.Area())
	println("re Primeter is", re.Primeter())
}

type Rectangle struct {
	length int
	width  int
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func (r *Rectangle) Primeter() int {
	return (r.length + r.width) * 2
}
