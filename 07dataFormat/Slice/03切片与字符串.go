package main

import "fmt"

func main03() {
	str := "hello,yvesyuan"
	a1 := str[0:6]
	fmt.Println(a1)
	//str[6] = '!'
	//a1[0] = '!'
}
func main031() {
	str := "hello,yvesyuan"
	a1 := []byte(str)
	a1[5] = '!'
	a1 = append(a1, '!')
	str = string(a1)
	fmt.Println(str)
}
func main() {
	str := "hello,yvesyuan"
	a1 := []rune(str)
	a1[5] = '好'
	a1 = append(a1, '!')
	str = string(a1)
	fmt.Println(str)
}
