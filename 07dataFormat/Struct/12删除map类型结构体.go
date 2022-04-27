package main

import "fmt"

type student12 struct {
	id   int
	name string
	age  int
}

func main() {
	ce := make(map[int]student12) //map类型的结构体	int->student12
	ce[1] = student12{1, "xiaolizi", 22}
	ce[2] = student12{2, "wang", 23}
	fmt.Println(ce)
	delete(ce, 2) //删除
	fmt.Println(ce)
}
