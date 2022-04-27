package main

import (
	"encoding/json"
	"fmt"
)

//`json:"skillName,omitempty"` 结构体Tag，omitempty过滤空值
type Skill struct {
	Name  string `json:"skillName,omitempty"` //钢琴 十级、小提琴 三级
	Leval int    `json:"skillLeval"`
}

type Actor struct { //个人信息
	Name   string `json:"actorName"`
	Age    int    `json:"actorAge"`
	Skills []Skill
}

func main() {
	yves := &Actor{
		Name: "yves",
		Age:  18,
		Skills: []Skill{
			{Name: "Piano", Leval: 10},
			{Name: "Violine", Leval: 3},
			{Name: "", Leval: 4},
		},
	}
	data, _ := json.Marshal(yves)
	fmt.Println(string(data))
	//{"Name":"yves","Age":18,"Skills":[{"Name":"Piano","Leval":10},{"Name":"Violine","Leval":3}]}
	//{"actorName":"yves","actorAge":18,"Skills":[{"skillName":"Piano","skillLeval":10},{"skillName":"Violine","skillLeval":3}]}
	back := &Actor{}
	json.Unmarshal(data, back)
	fmt.Println(back)
	//&{yves 18 [{Piano 10} {Violine 3} { 4}]}
}
