package main

import "fmt"

type student3 struct {
	name           string
	age, id, heavy int
}

func newStu3(age, id, heavy int, name string) *student3 { //构造函数
	return &student3{name: name, age: age, id: id, heavy: heavy}
}

//值类型
func (s student3) addAge() {
	s.age += 1
}

//指针类型
func (s *student3) addHeavy() {
	s.heavy += 10
}

func main() {
	myStu := newStu3(20, 01, 100, "yves")
	fmt.Println(myStu)
	myStu.addAge()
	myStu.addHeavy()
	fmt.Println(myStu)
}
