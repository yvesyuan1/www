package main

import "fmt"

func main() {
	//创建
	m1 := map[int]string{1: "yves", 2: "libing"}

	//赋值
	m1[3] = "jack"

	//删除
	delete(m1, 2)

	//查找元素，判断存在
	value, ok := m1[1]
	if ok {
		fmt.Println("m1[1]存在，为", value)
	} else {
		fmt.Println("m1[1]不存在")
	}

	//遍历
	for k, v := range m1 {
		fmt.Println(k, ":", v)
	}

}
