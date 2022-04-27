package main

import "fmt"

// 本函数测试入口参数和返回值情况
func dummy(b int) int {
	// 声明一个变量c并赋值
	var c int
	c = b
	return c
}

// 空函数, 什么也不做
func void() {
}

func main() {
	// 声明a变量并打印
	var a int
	// 调用void()函数
	void()
	// 打印a变量的值和dummy()函数返回
	fmt.Println(a, dummy(0))
}

//使用 go run 运行程序时，-gcflags 参数是编译参数。
//其中 -m 表示进行内存分配分析，-l 表示避免程序内联，也就是避免进行程序优化。
//命令行：go run -gcflags "-m -l" 05变量逃逸分析.go
/*
# command-line-arguments
./05变量逃逸分析.go:23:13: ... argument does not escape
./05变量逃逸分析.go:23:13: a escapes to heap	//a 逃逸到堆
./05变量逃逸分析.go:23:22: dummy01(0) escapes to heap
0 0
*/
