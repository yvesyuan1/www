package main

import (
	"fmt"
	"regexp"
)

func main34() {
	r1 := regexp.MustCompile("ab+c")                        //a b(一次或多次) c
	fmt.Println(r1.FindString("abxbababbbbbbbbbbbbcavvvc")) //寻找第一个匹配的字符串
}

func main() {
	r1 := regexp.MustCompile("a.+c")                         //a 任何单个字符(一次或多次) c
	fmt.Println(r1.FindString("<html>ahhhcabcajjjc</html>")) //ahhhcabcajjjc
}
