package main

import "fmt"

func main() {
	type Address struct {
		Province string
		City     string
		Street   string
	}

	type Vcard struct {
		Name     string
		Addr     *Address
		Birthday string
		Picture  string
	}

	claris := new(Vcard)
	addr := &Address{"CCC", "AAA", "SSS"}
	claris.Name = "J"
	claris.Addr = addr
	claris.Birthday = "1000-0-0"
	claris.Picture = "www.picture.com"

	fmt.Println(claris)
}
