package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	TestString := "袁驿牛老板"
	md5run := md5.New()
	md5run.Write([]byte(TestString))
	res := md5run.Sum(nil)
	fmt.Printf("%x\n", res)

	TestString2 := "袁驿小牛老板"
	md5run2 := md5.New()
	md5run2.Write([]byte(TestString2))
	res2 := md5run2.Sum(nil)
	fmt.Printf("%x\n", res2)
}
