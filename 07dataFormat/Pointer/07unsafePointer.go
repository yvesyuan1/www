package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 10
	var intP *int = &i
	fmt.Println(intP, *intP)
	var floatP *float64 = (*float64)(unsafe.Pointer(intP))
	//*floatP += 10 //*floatP的值为10，加法不能改变*floatP的值，i的值会错误4621819117588971520
	*floatP *= 10 //乘法能改变*floatP的值，为0，但i的值会正确变化，为100
	//todo 这里有Bug
	fmt.Printf("%p %f\n", floatP, *floatP)
	fmt.Printf("%d", i)
}

func main04() {
	arr := [3]int{1, 2, 3}
	//先将arr的首地址转为可以计算的uintptr类型，再转为unsafePointer，再转为*int
	p := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr)) + unsafe.Sizeof(arr[0])))
	*p = 5
	fmt.Println(arr)
}
