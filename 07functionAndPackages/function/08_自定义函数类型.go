package main

import "fmt"

func ADD(a, b int) int {
	return a + b
}
func SUB1(a, b int) int {
	return a - b
}

type MyFuncType func(int, int) int

func main() {
	var f1 MyFuncType = ADD
	fmt.Println(f1(1, 2))
	fmt.Printf("%T\t%V\n", f1, f1)
	f1 = SUB1
	fmt.Println(f1(1, 2))
	fmt.Printf("%T\t%V\n", f1, f1) //地址不同，类型相同
}
