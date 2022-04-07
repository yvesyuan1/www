package testPackage

import "fmt"

func init() {
	fmt.Println("加法初始化")
}

func Add(a, b int) int {
	return a + b
}

func AddString(a, b string) string {
	return a + b
}
