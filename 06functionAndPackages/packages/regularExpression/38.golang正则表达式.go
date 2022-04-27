package main

import (
	"fmt"
	"regexp"
)

func main38() {
	mystr := "my email is yvesyuan@foxmail.com" //校验
	matched, _ := regexp.MatchString(`[\w.%+-]+@[\w]+(?:\.[\w]+)+`, mystr)
	fmt.Printf("%v", matched)
}

func main39() {
	var str = `
lsiugaf
liurw;qfioub
skfjbsfqf
yves@qq.com
88827@163.com
sioufbqp*@&#@bb.com
souyuyu@sohu.com`

	re := regexp.MustCompile(`[\w.%+$-]+@[\w]+(?:\.[\w]+)+`)
	matches := re.FindAllStringSubmatch(str, -1) //n为检索次数，-1为不限
	for _, match := range matches {
		fmt.Println(match)
	}
}

func main() {
	mystr := "my email is yvesyuan@foxmail.com" //校验
	re := regexp.MustCompile(`[\w.%+$-]+@[\w]+(?:\.[\w]+)+`)
	replaceMystr := re.ReplaceAllString(mystr, "******")
	fmt.Printf("%s", replaceMystr) //my email is ******
}
