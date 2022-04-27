package main

/*
***重要***
 */

import "fmt"

type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student) { //ce本质是一个地址
	fmt.Printf("demo函数中，操作前：    ce的值为:%p\n", ce)
	ce[1].age = 999
	//有效    //ce的值是main函数中ce的地址，修改ce对应的数值，就会修改main函数里面的ce对应的数值
	ce = append(ce, student{3, "xiaowang", 56})
	//无效 	//这里append修改了ce值，也就是修改了ce所指向的内容，这时候无论做什么改变，都不会影响到main函数
	fmt.Printf("demo函数中，操作后：    ce的值为:%p\n", ce)
	ce[1].age = 99999
	//无效	//因为这里ce并不指向main函数里的ce，做什么都无效
}
func main() {
	var ce []student //定义一个切片类型的结构体
	ce = []student{
		student{1, "xiaoming", 22},
		student{2, "xiaozhang", 33},
	}
	fmt.Println(ce)
	fmt.Printf("main函数中，函数调用前：ce的值为:%p\n", ce)
	demo(ce)
	fmt.Printf("main函数中，函数调用后：ce的值为:%p\n", ce)
	fmt.Println(ce)
}

/*
	[{1 xiaoming 22} {2 xiaozhang 33}]
	main函数中，函数调用前：ce的值为:0x14000128040
	demo函数中，操作前：    ce的值为:0x14000128040
	demo函数中，操作后：    ce的值为:0x14000134000
	main函数中，函数调用后：ce的值为:0x14000128040
	[{1 xiaoming 22} {2 xiaozhang 999}]
*/
