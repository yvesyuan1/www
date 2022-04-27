package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := make(map[int]string, 5)
	map1[1] = "www.尹成微信18510341407"
	map1[2] = "rpc.尹成微信18510341407"
	map1[5] = "ceshi"
	map1[3] = "xiaohong"
	map1[4] = "xiaohuang"
	sli := []int{} //定义切片
	for k, _ := range map1 {
		sli = append(sli, k)
	}
	sort.Ints(sli)
	for i := 0; i < len(map1); i++ {
		fmt.Println(map1[sli[i]])
	}
}
