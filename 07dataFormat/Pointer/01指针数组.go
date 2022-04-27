package main

import "fmt"

func main() {
	a, b, c := 1, 2, 3
	pointers := []*int{&a, &b, &c}
	for _, p := range pointers {
		fmt.Println(*p)
	}
}
