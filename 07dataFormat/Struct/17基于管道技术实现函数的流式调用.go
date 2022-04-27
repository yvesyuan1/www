package main

import (
	"fmt"
)

type user01 struct {
	name string
	age  int
}

// interface{}作为返回值，可以返回任意类型
func FilterAge(users []user01) interface{} {
	var slice []user01
	for _, u := range users {
		if u.age >= 18 && u.age <= 35 {
			slice = append(slice, u)
		}
	}
	return slice
}

//将年龄保存为数组
func MapAgeToSlice(users []user01) interface{} {
	var slice []int
	for _, u := range users {
		slice = append(slice, u.age)
	}
	return slice
}

//年龄的累和
func SumAge(users []user01, pipes ...func([]user01) interface{}) int { //pipes 一系列的操作函数
	var sum = 0
	var ages []int
	for _, pipe := range pipes {
		result := pipe(users)
		//result.(type)	判断result的类型
		switch result.(type) {
		case []user01: //这里是方便执行FilterAge函数
			users = result.([]user01) //类型转换？
		case []int:
			ages = result.([]int) //这里是方便执行MapAgeToSlice函数
		}
	}
	if len(ages) == 0 {
		fmt.Println("管道没有提取年龄")
		return 0
	} else {
		for _, age := range ages {
			sum += age
		}
	}
	return sum
}

func main() {
	var users = []user01{
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
	fmt.Println(SumAge(users, FilterAge, MapAgeToSlice))
}
