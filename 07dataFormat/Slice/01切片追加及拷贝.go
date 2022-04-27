package main

import "fmt"

func main01() {
	var a = [...]int{1, 2, 3, 4, 5, 6, 7}
	var slice1 = a[0:3:4]
	fmt.Println(slice1, len(slice1), cap(slice1))

	slice1 = append(slice1, 8, 9)
	fmt.Println(slice1, len(slice1), cap(slice1))

}

func main() {
	var a = []int{1, 2, 3}
	var b = []int{11, 22, 33}
	a = append(a, b...)              //s如果使用s...符号解压缩切片，则可以将切片直接传递给可变参数函数。在这种情况下，不会创建新的切片。
	a = append([]int{0, 0, 0}, a...) //在前面追加切片
	fmt.Println(a)
}
