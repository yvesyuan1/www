package main

import "fmt"

//类型定义
type NewInt int

//类型别名
type myInt = int

//结构体定义
type myStuct struct {
	name string
	age  int
	sex  string
}

//基本实例化
func main010() {
	var s1 myStuct
	s1.name = "yves"
	s1.age = 16
	s1.sex = "male"
	fmt.Println(s1)
}

//匿名结构体
func main011() {
	var user struct {
		Name string
		Age  int
	} //注意是分号
	user.Name = "yves"
	user.Age = 16
	fmt.Println(user)
}

//创建指针类型结构体
func main012() {
	var s2 = new(myStuct)
	s2.name = "yves"
	s2.age = 16
	s2.sex = "male"
	fmt.Println(s2)
}

//取结构体的地址实例化
func main013() {
	var s3 = &myStuct{}
	s3.sex = "make"
}

//使用键值对初始化
func main014() {
	s4 := myStuct{
		name: "yves", //注意逗号
		age:  14,
	}
	fmt.Println(s4)
}

//对结构体指针进行键值对初始化
func main() {
	s5 := &myStuct{
		"yves", //注意逗号
		14,
		"male",
	}
	fmt.Println(s5)
}
