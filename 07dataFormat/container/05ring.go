package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(10)    //返回第一个元素的地址
	fmt.Println(r.Len()) //长度
	fmt.Println(r)

	for i := 1; i < 11; i++ {
		r.Value = i
		r = r.Next()
	} //最后r回到第一位
	fmt.Println("*********************")
	r.Do(func(i interface{}) { //此时r在第一位
		fmt.Println(i)
	}) //打印

	fmt.Println("*********************")
	fmt.Println(r)
	r = r.Next() //向后一位
	fmt.Println(r)
	r = r.Move(3) //向后三位
	fmt.Println(r)
	r = r.Move(-3) //向前三位
	fmt.Println(r)
	r = r.Prev() //向前一位
	fmt.Println(r)

	r5 := r.Move(4) //从第一位向右走4位 到第五位
	r8 := r5.Move(3)
	rno := r5.Link(r8) //把第8位连接到第5位	//也就是删除了6，7位	//返回第5位的后一位地址
	fmt.Println("*********************")
	fmt.Println(rno)
	r.Do(func(i interface{}) { //此时r在第一位
		fmt.Println(i)
	}) //打印

}
