package main

import (
	"fmt"
	"strings"
)

func stringProcess(list []string, chain []func(string) string) {
	for index, str := range list {
		result := str
		for _, proc := range chain {
			result = proc(result)
		}
		list[index] = result
	}
}

// 自定义的移除前缀的处理函数
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func main() {
	// 待处理的字符串列表
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	// 使用的函数列表
	chain := []func(string) string{
		removePrefix,      //去掉前缀
		strings.TrimSpace, //去空格
		strings.ToUpper,   //转大写
	}

	// 处理字符串
	stringProcess(list, chain)

	// 输出处理好的字符串
	for _, str := range list {
		fmt.Println(str)
	}
}
