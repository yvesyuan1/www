package main

import (
	"fmt"
	"regexp"
)

func main37() {
	r1 := regexp.MustCompile(`Mobile:([0-9]{11})`)
	fmt.Println(r1.FindString("email:yves_yuan@foxmail.com Mobile:13199999999"))

}
func main() {
	r1 := regexp.MustCompile(`(Mobile):([0-9]{11})`)
	fmt.Println(r1.FindString("email:yves_yuan@foxmail.com Mobile:13199999999"))

}
