package main

import "fmt"

type student struct {
	name string
	age  int
}

func change(stu *student) {
	stu.age = 24
}

func main() {
	stu := new(student) //stu为结构体指针	//内存分配在堆上
	stu.name = "yves"
	change(stu)
	fmt.Println(stu)
}

func main02() {
	var stu *student = &student{ // 定义结构体指针 = 结构体地址  //内存分配在堆上
		name: "yves",
		age:  20,
	}
	stu.name = "BossNiu"
	change(stu)
	fmt.Println(stu)
}

func main01() {
	var stu1 student = student{
		name: "test",
		age:  2,
	}
	stu1.name = "test01"
	change(&stu1)
	fmt.Println(stu1)

}
