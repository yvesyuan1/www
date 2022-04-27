package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InitArray(num []int) { //生成随机数
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(num); i++ {
		num[i] = rand.Intn(1000)
	}
}

func sortArray(num []int) { //冒泡
	for j := 0; j < len(num)-1; j++ {
		for i := 0; i < len(num)-1-j; i++ {
			if num[i] < num[i+1] {
				num[i], num[i+1] = num[i+1], num[i]
			}
		}
	}
}

func main() {
	s := make([]int, 1000, 1000)
	InitArray(s)
	fmt.Println(s)
	sortArray(s)
	fmt.Println(s)
}
