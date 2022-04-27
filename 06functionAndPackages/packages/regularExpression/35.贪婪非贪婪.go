package main

import (
	"fmt"
	"regexp"
)

//+尽量多的匹配、？匹配到一个就停
func main35() {
	r1 := regexp.MustCompile("<.+>")                         //贪婪： <任何字符(一次或多次)>
	fmt.Println(r1.FindString("<html>ahhhcabcajjjc</html>")) //<html>
}

func main() {
	r1 := regexp.MustCompile("<.+?>")                        //贪婪： <任意字符(零次或一次)>
	fmt.Println(r1.FindString("<html>ahhhcabcajjjc</html>")) //<html>
}
