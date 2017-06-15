package sort

type Person struct {
	firstName string
	lastName string
}

type Persons []*Person

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int) 
}

func (p Persons) Len() int {
	return len(p)
}

func (p Persons) Less(i, j int) bool {
	return (p[i].firstName + p[i].lastName) < (p[j].firstName + p[j].lastName) 
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}




