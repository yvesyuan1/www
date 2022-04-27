package main

import "fmt"

func myAppend(slice []int) []int {
	//这个会改变外面s的值
	fmt.Printf("%p\n", slice)
	slice[0] = 0
	// 这里增加了slice长度但不会改变外面s的值
	fmt.Printf("%p\n", slice)
	slice = append(slice, 100)
	fmt.Printf("%p\n", slice)
	return slice
}
func myAppendPtr(slice *[]int) {
	// 会改变外层 s 本身
	*slice = append(*slice, 100)
	return
}
func main() {
	s := []int{1, 1, 1}
	newS := myAppend(s)
	fmt.Println(s)
	fmt.Println(newS)
	s = newS
	myAppendPtr(&s)
	fmt.Println(s)
}
