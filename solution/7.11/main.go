package main

import "fmt"

func main() {
	des := []int{1,2,3,4,5,6,7}
	tar := []int{9,9,9}
	fmt.Println(InsertStringSlice(des, tar, 3))
}

func InsertStringSlice(des, tar []int, index int) []int {
	dl := len(des)
	tl := len(tar)
	ns := make([]int, dl + tl)
	copy(ns, des[:index])
	copy(ns[index:], tar)
	copy(ns[(index+tl):], des[index:])
	return ns
}
