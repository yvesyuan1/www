package main

import (
	"GoFullStackDeveloper/06functionAndPackages/packages/testPackage"
	"fmt"
)

func main() {
	fmt.Println(testPackage.Add(1, 2))
	fmt.Println(testPackage.AddString("hello", "world"))
}
