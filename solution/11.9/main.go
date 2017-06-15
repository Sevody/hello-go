package main
import "fmt"

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}
func (p Person) ShowP() string {
	return "get in function"
}
type Any interface{}

func main() {
	var val Any
	val = 5
	fmt.Printf("val has the value: %v\n", val)
	val = str
	fmt.Printf("val has the value: %v\n", val)
	pers1 := new(Person)
	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1
	fmt.Printf("val has the value: %v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T and value is %v\n", t, t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T and vlue is %s\n", t, t.ShowP())
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}