package main

import (
	"flag"
	"fmt"
)

var Input_pstrName = flag.String("name", "gerry", "input ur name")
var Input_piAge = flag.Int("age", 20, "input ur age")
var Input_flagvar int

func Init() {
	flag.IntVar(&Input_flagvar, "flagname", 1234, "help message for flagname")
}

func main() {
	Init()
	flag.Parse() //对命令行参数进行解析

	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg()) //命令行参数后的其他参数、参数的个数
	for i := 0; i != flag.NArg(); i++ {                       //遍历每一个参数
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println("name=", *Input_pstrName)
	fmt.Println("age=", *Input_piAge)
	fmt.Println("flagname=", Input_flagvar)
}
