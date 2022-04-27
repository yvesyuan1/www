package main

import "fmt"

func swap(s *[]int) {
	l := len(*s)
	for i := 0; i < l/2; i++ {
		(*s)[i], (*s)[l-i-1] = (*s)[l-i-1], (*s)[i]
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	swap(&s)
	fmt.Println(s)
}
