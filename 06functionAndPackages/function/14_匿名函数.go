package main

import (
	"fmt"
	"time"
)

func main1() {
	func(data int) { fmt.Println("hello", data) }(1024)
	f1 := func(data float64) { fmt.Println("hello", data) }
	f1(12.34)
}

func useTimes(myfunc func(int) int, arg int) {
	starttime := time.Now() //开始时间
	myfunc(arg)             //调用函数
	fmt.Println(time.Since(starttime))
} //计算函数运行时长

func main2() {
	useTimes(func(arg int) (sum int) {
		for i := 0; i < arg; i++ {
			sum += i
		}
		return sum
	}, 10000000000)
}

func squres() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func main() {
	f := squres() //f存放的是squres返回的函数的地址
	//fmt.Println(f)	//0x102727b30
	fmt.Println(f()) //调用f（），就是调用返回的那个函数，每次调用x都会++
	fmt.Println(f())
	fmt.Println(f())
}
