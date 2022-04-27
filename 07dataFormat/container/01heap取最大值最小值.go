package main

import (
	"container/heap"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type myHeap []int

//在go中，并不像java提供一个可以直接使用的容器类，
//而是提供一个接口，需要你实现这个接口里的方法来自定义一个堆/集合

//比大小
func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

//排序
//个人理解：干掉最后一个元素，并返回最后一个元素
func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}

//压入
func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

//删除
func (h *myHeap) Remove(idx int) interface{} {
	h.Swap(idx, h.Len()-1) //把要干的放到最后一位，用pop函数给他干了
	return h.Pop()
}

func main() {
	h := new(myHeap)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		h.Push(rand.Intn(1000))
	}
	fmt.Println(h)

	for i := 0; i < 10; i++ {
		//对heap进行初始化，形成小（大）根堆，可以理解为排序
		//最大堆：根结点的键值是所有堆结点键值中最大者。
		//最小堆：根结点的键值是所有堆结点键值中最小者。
		heap.Init(h)
		fmt.Println(h)
		//testPrintHeap(h)

		h.Swap(0, h.Len()-1) //把根节点与最后一个叶子节点交换
		fmt.Println(h.Pop()) //取出最后一个叶子节点即为最小值
	}

}

func printHeap(h *myHeap) {
	deep := math.Floor(math.Log2(float64(h.Len()-1))) + 1
	for i := 1.0; i <= deep; i++ {
		for j := 1; j <= int(math.Pow(2.0, i)); j++ {
			num := int(math.Pow(2.0, i-1)) + j
			fmt.Printf("%d\t", (*h)[num])
		}
	}
}
