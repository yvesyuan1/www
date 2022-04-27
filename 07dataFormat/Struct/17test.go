package main

import (
	"fmt"
	"strings"
)

type user struct {
	name string
	age  int
}

//使用一条语句一个函数
//实现对users结构体数组"名字不含四"，"年龄大于18小于35"的筛选，并求年龄和

func filterName(users []user) interface{} {
	var slice []user
	for _, u := range users {
		if !strings.Contains(u.name, "四") {
			slice = append(slice, u)
		}
	}
	return slice
}

func filterAge(users []user) interface{} {
	var slice []user
	for _, u := range users {
		if u.age >= 18 && u.age <= 35 {
			slice = append(slice, u)
		}
	}
	return slice
}

func ToAgeSlice(users []user) interface{} {
	var age []int
	for _, u := range users {
		age = append(age, u.age)
	}
	return age
}

func getAgeSum(users []user, pipe ...func(users []user) interface{}) int {
	var ageSum int
	var ages []int
	for _, f := range pipe {
		result := f(users)
		switch result.(type) {
		case []int:
			ages = result.([]int)
		case []user:
			users = result.([]user)
		}
	}
	for _, age := range ages {
		ageSum += age
	}
	//fmt.Println(users)
	return ageSum
}

func main() {
	var users = []user{
		{
			name: "张三",
			age:  18,
		},
		{
			name: "李四",
			age:  22,
		},
		{
			name: "王五",
			age:  20,
		},
		{
			name: "赵六",
			age:  -10,
		},
		{
			name: "孙七",
			age:  60,
		},
		{
			name: "周八",
			age:  10,
		},
	}
	fmt.Println(getAgeSum(users, filterName, filterAge, ToAgeSlice))
}
