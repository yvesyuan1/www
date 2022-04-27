package main

import "fmt"

func changeMap(m map[int]string, n int) {
	m[n] = "changed"
}

func buildMap() map[int]string {
	m1 := map[int]string{1: "yves", 2: "libing", 10: "hello"}
	return m1
}

func main03() {
	m1 := map[int]string{1: "yves", 2: "libing", 10: "hello"}
	changeMap(m1, 2)
	fmt.Println(m1)
}

func main() {
	m := buildMap()
	fmt.Println(m)
}
