package main

import "fmt"

func main() {
	//s1 := make([]int, 5, 10)
	//fmt.Println(len(s1), cap(s1))
	//
	//var array1 = [...]int{1, 2, 3, 4, 5}
	//s2 := array1[0:4]
	//fmt.Println(s2)
	//fmt.Printf("%p", s2) //s2地址

	data := [][]int{[]int{1, 2, 3}, []int{4, 5, 6, 7}, []int{8}}
	fmt.Println(data)
}
