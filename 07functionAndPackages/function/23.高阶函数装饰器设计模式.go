package main

import (
	"fmt"
	"time"
)

func multiply(a, b int) int {
	return a * b
}

func useTimes2(myfunc func(a, b int) int, a, b int) int { //高级函数
	starttime := time.Now() //开始时间
	c := myfunc(a, b)       //调用函数
	fmt.Println(time.Since(starttime))
	return c
} //计算函数运行时长

func main() {
	a := 2
	b := 8
	c := useTimes2(multiply, a, b)
	fmt.Printf("%d x %d = %d\n", a, b, c)
}
