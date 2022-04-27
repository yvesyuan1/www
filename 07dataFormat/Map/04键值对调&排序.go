package main

import (
	"fmt"
	"sort"
)

func main04() {
	//键值对调
	m1 := map[int]string{1: "yves", 2: "libing", 10: "hello"}

	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)
}

func main() {
	//排序
	m1 := map[int]string{1: "yves", 20: "libing", 10: "hello"}

	keys := make([]int, 0)
	for k, _ := range m1 {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println(m1[k])
	}

}
