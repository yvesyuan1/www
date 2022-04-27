package main

import "fmt"

type student2 struct {
	name           string
	age, id, heavy int
}

func newStu2(age, id, heavy int, name string) *student2 { //构造函数
	return &student2{name: name, age: age, id: id, heavy: heavy}
}

func (s student2) studying() {
	fmt.Println(s.name, "正在学习")
}

func main() {
	myStu := newStu2(20, 01, 100, "yves")
	fmt.Println(myStu)
	myStu.studying()
}
