package main

import (
	"math/rand"
	"strconv"
	"fmt"
)

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Float64Array []float64

func (f Float64Array) Len() int {
	return len(f)
}

func (f Float64Array) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f Float64Array) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f Float64Array) List() string {
	str := ""
	for i, v := range f {
		str += "[" + strconv.Itoa(i) + ":" + strconv.FormatFloat(v, 'f', 2, 64) + "]"
	}
	return str
}

func (f Float64Array) String() string {
	return f.List()
}
func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}
func NewFloat64Array(l int) Float64Array {
	return make([]float64, l)
}
func Fill() Float64Array {
	f := NewFloat64Array(10)
	for i := 0; i < f.Len(); i++ {
		f[i] = rand.Float64() * 100.00
	}
	return f
}

func main() {
	f := Fill()
	Sort(f)
	fmt.Println(f)
}