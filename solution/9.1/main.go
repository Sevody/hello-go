package main

import (
	"container/list"
	"fmt"
)

func main() {
	ls := list.New()
	ls.PushBack(101)
	ls.PushBack(102)
	ls.PushBack(103)

	for e := ls.Front(); e != nil; e = e.Next() {
		fmt.Println(e)
	}
}
