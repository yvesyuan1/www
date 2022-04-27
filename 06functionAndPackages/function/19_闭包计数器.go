package main

import "fmt"

func main3() {
	pos := adder()
	for i := 0; i < 101; i++ {
		fmt.Printf("i=%d\t", i)
		fmt.Println(pos(i))

	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		fmt.Printf("sum up=%d\t", sum)
		sum += x
		fmt.Printf("sum down=%d\t", sum)
		return sum
	}
}

//计数器,闭包
func Counter() func() int {
	i := 0
	res := func() int {
		i += 1
		return i
	}
	fmt.Println("内部函数res", res)
	return res
}

func main() {
	myfunc := Counter()
	fmt.Println(myfunc())
	fmt.Println(myfunc())
	fmt.Println(myfunc())
	fmt.Println(myfunc())

	myfunc1 := Counter()
	fmt.Println(myfunc1())
	fmt.Println(myfunc1())
	fmt.Println(myfunc1())
	fmt.Println(myfunc1())
}
