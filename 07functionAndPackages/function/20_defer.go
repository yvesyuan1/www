package main

import "fmt"

//延迟执行
func main12() {
	defer fmt.Println("hello1") //压入栈
	defer fmt.Println("hello2")
	defer fmt.Println("hello3")
	fmt.Println("hello4") //先执行这句，再执行defer
	//没有defer自上而下，有defer自下而上
}

func main22() {
	defer func() {
		fmt.Println("hello123")
	}()
	fmt.Println("hello yves")
}

func main33() {
	var a, b = 10, 20
	defer func(a, b int) { fmt.Println("defer", a, b) }(a, b) //虽然函数被压入栈，但赋值已经完成，所以ab为10，20
	a = 100
	b = 200
	fmt.Println(a, b)
}

func f1() (r int) {
	defer func() { r++ }() //step2 r++ r = 1
	r = 0                  //step1 r = 0
	return                 //step3 r = 1
}
func main44() {
	i := f1()
	fmt.Println(i)
}

func double(x int) int {
	return 2 * x
}
func triple(x int) (r int) {
	defer func() { r += x }() //step2 r = 3+6 	r = 9
	return double(x)          //step1 double(3) r=6		//step3 return 9
}
func main() {
	fmt.Println(triple(3))
}
