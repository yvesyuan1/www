package main

import (
	"container/list"
	"fmt"
	"time"
)

func testSlice() {
	start := time.Now()
	_slice := make([]int, 1)
	for i := 0; i < 100000000; i++ {
		_slice = append(_slice, i)
	}
	fmt.Println(time.Now().Sub(start).String())
}

func testListToBack() {
	start := time.Now()
	_list := list.New()
	for i := 0; i < 100000000; i++ {
		_list.PushBack(i)
	}
	fmt.Println(time.Now().Sub(start).String())
}

func testListToFront() {
	start := time.Now()
	_list := list.New()
	for i := 0; i < 100000000; i++ {
		_list.PushFront(i)
	}
	fmt.Println(time.Now().Sub(start).String())
}

func main() {
	testSlice()
	testListToBack()
	testListToFront()
}
