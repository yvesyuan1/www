package main

import "fmt"

func main() {
	type test struct {
		a int8
		b int8
		c int16
		d int32
		e int64
	}
	n := test{
		1, 2, 3, 4, 5,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
	fmt.Printf("n.d %p\n", &n.e)
}
