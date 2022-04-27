package main

import "fmt"

func myAppend01(slice01 []int) []int {
	slice01[2] = 1111
	slice01 = append(slice01, 100)
	return slice01
}

func myAppend02(slice02 []int) {
	slice02[2] = 1111
	slice02 = append(slice02, 100)
}

func main() {
	s := make([]int, 3, 5)
	s[2] = 8888
	myAppend02(s)
	fmt.Printf("%p  ", s)
	fmt.Println(s)
}
