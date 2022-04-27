package main

import (
	"encoding/json"
	"fmt"
)

type Student1 struct {
	Id     int
	Gender string
	Name   string
}

type Class1 struct {
	Title   string
	Student []*Student1 //数组
}

func main() {
	jsonStr := `{"Title":"软工2班","Student":[{"Id":0,"Gender":"male","Name":"stu0"},{"Id":1,"Gender":"male","Name":"stu1"},{"Id":2,"Gender":"male","Name":"stu2"},{"Id":3,"Gender":"male","Name":"stu3"},{"Id":4,"Gender":"male","Name":"stu4"},{"Id":5,"Gender":"male","Name":"stu5"},{"Id":6,"Gender":"male","Name":"stu6"},{"Id":7,"Gender":"male","Name":"stu7"},{"Id":8,"Gender":"male","Name":"stu8"},{"Id":9,"Gender":"male","Name":"stu9"},{"Id":10,"Gender":"male","Name":"stu10"},{"Id":11,"Gender":"male","Name":"stu11"},{"Id":12,"Gender":"male","Name":"stu12"},{"Id":13,"Gender":"male","Name":"stu13"},{"Id":14,"Gender":"male","Name":"stu14"},{"Id":15,"Gender":"male","Name":"stu15"},{"Id":16,"Gender":"male","Name":"stu16"},{"Id":17,"Gender":"male","Name":"stu17"},{"Id":18,"Gender":"male","Name":"stu18"},{"Id":19,"Gender":"male","Name":"stu19"}]}
`
	c1 := &Class1{}
	err := json.Unmarshal([]byte(jsonStr), c1)
	if err != nil {
		return
	}
	fmt.Println(c1)
}
