package main

import (
	"fmt"
)

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Println(a.name, "在走路")
}

type Dog struct {
	age     int
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) say() {
	fmt.Println("汪汪汪")
}

func (d *Dog) move() { //重写
	fmt.Println(d.name, "在跑")
}

func main() {
	d := &Dog{
		age:    2,
		Animal: &Animal{name: "dcy"},
	}
	d.move()
	d.say()
	d.Animal.move()
}
