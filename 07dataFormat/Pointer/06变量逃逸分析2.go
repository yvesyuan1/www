package main

import "fmt"

// 声明空结构体测试结构体逃逸情况
type Data struct {
}

func dummy01() *Data {
	// 实例化c为Data类型
	var c Data
	//返回函数局部变量地址
	return &c //c作为局部变量，但还是会被引用，所以这个时候不能放在栈上
}
func main() {
	fmt.Println(dummy01())
}

/*
# command-line-arguments
./06变量逃逸分析2.go:11:6: moved to heap: c //将 c 移到堆中
./06变量逃逸分析2.go:16:13: ... argument does not escape
&{}
*/
