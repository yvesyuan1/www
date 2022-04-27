package main

import "fmt"

func SUB(a, b int) int {
	return a - b
}

//函数本质是地址

func main() {
	fmt.Printf("%T\n", SUB)

	ak47 := SUB
	fmt.Printf("%T\n", ak47)
	fmt.Println(ak47(100, 30))

	fmt.Println(SUB)
	fmt.Println(ak47) //地址相同
}
