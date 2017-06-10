package main

func main() {
	var arr [15]int
	for i:=0;i<len(arr);i++{
		arr[i] = i*2
	}
	for _,v := range arr{
		println(v)
	} 
}