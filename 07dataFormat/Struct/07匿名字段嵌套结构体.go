package main

import "fmt"

type Location struct {
	address   string
	creatTime string
}

type Email struct {
	address   string
	creatTime string
}

type User struct {
	id       int
	name     string
	email    Email
	Location //匿名字段
}

func main() {
	var user User
	user.id = 10
	user.name = "yves"
	user.Location.address = "NUIST"
	user.Location.creatTime = "2021"
	user.email.address = "yves_yuan@foxmail.com"
	user.email.creatTime = "2018"
	fmt.Println(user)
}
