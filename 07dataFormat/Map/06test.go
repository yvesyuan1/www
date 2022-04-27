package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Map-Reduce_Filter设计模式
/* test
1.从users的map中筛选出名字里不含"六"的users
2.计算剩下所有人的年龄和
3.使用Map_Reduce_Filter设计模式
*/

//第一步，过滤器
func itemFilter(items []map[string]string, f func(map[string]string) bool) []map[string]string {
	newMapSlice := make([]map[string]string, len(items))
	for _, item := range items {
		if f(item) {
			newMapSlice = append(newMapSlice, item)
		}
	}
	return newMapSlice
}

//第二步，将map转化成字符串切片
func mapToString(items []map[string]string, f func(map[string]string) string) []string {
	newSlice := make([]string, len(items))
	for _, item := range items {
		newSlice = append(newSlice, f(item))
	}
	return newSlice
}

//第三部，reduce求和函数
func getAgeSum(items []string, f func(string) int) int {
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

	//筛去名字里面含"六"的user
	newUsers := itemFilter(users, func(user map[string]string) bool {
		if strings.Contains(user["name"], "六") {
			return false
		} else {
			return true
		}
	})

	//创造age字符串切片
	sliceOfAge := mapToString(newUsers, func(user map[string]string) string {
		return user["age"]
	})

	//获得年龄和
	ageSum := getAgeSum(sliceOfAge, func(s string) int {
		age, _ := strconv.Atoi(s)
		return age
	})

	fmt.Println(ageSum)
}
