package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("now: %v;\tnow type:%T\n", now, now)
	fmt.Println(now.Month())
	fmt.Println(int(now.Month()))
	fmt.Println(now.Weekday())

	//时间格式化
	fmt.Printf("当前：%d-%d-%d %d %d:%d\n", now.Year(), now.Month(), now.Day(), now.Weekday(), now.Hour(), now.Minute())
	fmt.Printf(now.Format("2006-1-2 15:04:05")) //标砖案例
}
