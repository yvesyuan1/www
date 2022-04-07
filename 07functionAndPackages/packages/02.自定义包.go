package main

import (
	"GoFullStackDeveloper/07functionAndPackages/packages/testPackage"
	"fmt"
)

func main() {
	fmt.Println(testPackage.Add(1, 2))
	fmt.Println(testPackage.AddString("hello", "world"))
}
