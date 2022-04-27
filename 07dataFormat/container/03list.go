package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack("yuan")
	l.PushFront("Hello,")
	element := l.PushBack("!")
	l.InsertAfter("01", element)
	l.InsertBefore("yves", element)
	l.Remove(element.Next())
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
