package main

import "fmt"

type student1 struct {
	name           string
	age, id, heavy int
}

func newStu(age, id, heavy int, name string) *student1 { //构造函数
	return &student1{name: name, age: age, id: id, heavy: heavy}
}

func main() {
	myStu := newStu(20, 01, 100, "yves")
	fmt.Println(myStu)
}
