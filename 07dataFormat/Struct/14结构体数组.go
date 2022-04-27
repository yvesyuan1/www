package main

import "fmt"

type Student14 struct {
	name  string
	id    int
	score int
}

func main() {
	students := []Student14{
		{
			name:  "yves",
			id:    0,
			score: 20,
		},
		{
			name:  "lee",
			id:    01,
			score: 99,
		},
		{
			name:  "hellen",
			id:    2,
			score: 50,
		},
	}
	sum := 0
	for _, stu := range students {
		sum += stu.score
	}
	fmt.Println(sum)
}
