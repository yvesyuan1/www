package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type myHeap1 []int

//在go中，并不像java提供一个可以直接使用的容器类，
//而是提供一个接口，需要你实现这个接口里的方法来自定义一个堆/集合

func (h *myHeap1) Len() int {
	return len(*h)
}

//压入
func (h *myHeap1) Push(v interface{}) {
	*h = append(*h, v.(int))
}

//实现二叉树的打印
func testPrintHeap(h *myHeap1) {
	deep := math.Floor(math.Log2(float64(h.Len()))) + 1
	for i := 1.0; i <= deep; i++ {
		for j := 1; j <= int(math.Pow(2.0, i-1)); j++ {
			num := int(math.Pow(2.0, i-1)) + j - 1
			if num > (*h).Len() {
				return
			}
			fmt.Printf("%d\t", (*h)[num-1])
		}
		fmt.Printf("\n")
	}
}

func main() {
	h := new(myHeap1)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		h.Push(rand.Intn(1000))
	}
	fmt.Println(h)
	testPrintHeap(h)

}
