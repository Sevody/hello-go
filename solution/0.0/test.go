package main

import (
	"fmt"
)


type Idol struct {
  Name string
  Age int
}

func (idol *Idol) sayHello() {
  fmt.Printf("私は%s, %d歳です。", idol.Name, idol.Age)
}

func main() {
  kotori := &Idol{Name:"南小鳥", Age: 17}
	kotori.sayHello()
}