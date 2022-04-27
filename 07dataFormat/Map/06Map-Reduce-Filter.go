package main

import (
	"fmt"
	"strconv"
)

//将字典类型切片转化为一个字符串类型切片
func mapToStrings(items []map[string]string, f func(map[string]string) string) []string {
	newSlice := make([]string, len(items)) //开辟切片
	for _, item := range items {
		newSlice = append(newSlice, f(item))
	}
	return newSlice
}

// 求和函数：
func fieldSum(items []string, f func(string) int) int {
	var sum int
	for _, item := range items {
		sum += f(item)
	}
	return sum
}

func main() {
	var users = []map[string]string{
		{
			"name": "张三",
			"age":  "18",
		},
		{
			"name": "李四",
			"age":  "22",
		},
		{
			"name": "王五",
			"age":  "20",
		},
		{
			"name": "赵六",
			"age":  "-10",
		},
		{
			"name": "孙七",
			"age":  "60",
		},
		{
			"name": "周八",
			"age":  "10",
		},
	}

	ageSlice := mapToStrings(users, func(user map[string]string) string { return user["age"] })

	sum := fieldSum(ageSlice, func(age string) int {
		intAge, _ := strconv.Atoi(age)
		return intAge
	})

	fmt.Println(sum)
}
