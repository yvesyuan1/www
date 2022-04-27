package main

import "fmt"

//生成器
func playerGen(name string) func() (string, int) {
	hp := 150
	return func() (string, int) { //返回一个闭包
		return name, hp
	}
}

func main() {
	gen := playerGen("xiaoMing")
	name, hp := gen()
	fmt.Println(name, hp)
}
