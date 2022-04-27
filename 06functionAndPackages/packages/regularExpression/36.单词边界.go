package main

import (
	"fmt"
	"regexp"
)

//注意这里使用的是`字符，esc下面那个
func main36() {
	r1 := regexp.MustCompile(`\bhel`)               //匹配一个单词边界
	fmt.Println(r1.FindString("hellooooo world"))   //hel
	fmt.Println(r1.FindString("aahellooooo world")) //空
	r2 := regexp.MustCompile(`rld\b`)               //匹配一个单词边界
	fmt.Println(r2.FindString("hellooooo world"))   //rld
	fmt.Println(r2.FindString("hellooooo worldee")) //空
}

func main() {
	r1 := regexp.MustCompile(`\Bhel`)               //匹配一个非单词边界
	fmt.Println(r1.FindString("hellooooo world"))   //空
	fmt.Println(r1.FindString("aahellooooo world")) //hel
	r2 := regexp.MustCompile(`rld\B`)               //匹配一个非单词边界
	fmt.Println(r2.FindString("hellooooo world"))   //空
	fmt.Println(r2.FindString("hellooooo worldee")) //rld
}
