package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id     int
	Gender string
	Name   string
}

type Class struct {
	Title   string
	Student []*Student //数组
}

func main() {
	//创建班级
	c := &Class{Title: "软工2班", Student: make([]*Student, 0, 200)}

	//追加20个学生
	for i := 0; i < 20; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%d", i), //字符串格式化
			Gender: "male",
			Id:     i,
		}
		c.Student = append(c.Student, stu)
	}

	//将结构体序列化为Json
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Printf("%s", data)
}
