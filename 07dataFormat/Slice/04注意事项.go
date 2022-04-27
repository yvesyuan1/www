package main

import "fmt"

func change01(slice01 []int) { //地址传递
	fmt.Printf("%p\n", slice01)
	slice01[0] = 0
}

func change02(slice01 []int) []int { //地址传递
	fmt.Printf("%p\n", slice01)
	slice01[0] = 0
	return slice01
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	change01(s1)
	fmt.Printf("%p\n", s1)
	fmt.Println(s1)

	change02(s1)
	fmt.Printf("%p\n", s1)
	fmt.Println(s1)

}
