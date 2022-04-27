package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

type Girl struct {
	name   string
	age    int
	tall   int
	beauty int
}

type GilrHeap []Girl

func (h *GilrHeap) Less(i, j int) bool {
	return (*h)[i].beauty > (*h)[j].beauty
}

func (h *GilrHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *GilrHeap) Len() int {
	return len(*h)
}

func (h *GilrHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}

func (h *GilrHeap) Push(v interface{}) {
	*h = append(*h, v.(Girl))
}

func (h *GilrHeap) Remove(idx int) interface{} {
	h.Swap(idx, h.Len()-1) //把要干的放到最后一位，用pop函数给他干了
	return h.Pop()
}

func main() {
	girls := new(GilrHeap)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		girls.Push(Girl{
			name:   fmt.Sprintf("妹子%d", i),
			age:    rand.Intn(30) + 18,
			tall:   rand.Intn(40) + 150,
			beauty: rand.Intn(100),
		})
	}
	fmt.Println(girls)

	for i := 0; i < 10; i++ {
		heap.Init(girls)
		fmt.Println(girls)

		girls.Swap(0, girls.Len()-1)
		fmt.Println(girls.Pop())
	}
}
