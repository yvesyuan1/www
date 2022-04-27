package main

import "fmt"

type student01 struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student01)
	stus := []student01{
		{name: "1cxy.net", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}
	/*使用forrange，每次产生的 key 和 value 其实是对应的 stus 里面值的拷贝，不是对应的 stus 里面的值的引用
	stu 是 stus 在for循环中申请的一个局部变量，每次循环都会拷贝 stus 中对应的值拷贝到stu
	迭代遍历之后，stu 每次会被重新赋值，而这个 map 中记录的 value 只不过是 stu 的内存地址。*/
	//for _, stu := range stus {
	//	fmt.Printf("%p\n", &stu)
	//	m[stu.name] = &stu
	//}

	//解决方法，每次重新申请一个变量
	for _, stu := range stus {
		s := stu
		fmt.Printf("%p\n", &s)
		m[stu.name] = &s
	}
	//使用fori的方法也可以
	//for i := 0; i < len(stus); i++ {
	//	m[stus[i].name] = &stus[i]
	//}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

}
