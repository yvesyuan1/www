package main

import (
	"container/ring"
	"fmt"
)

type gamePlayer struct {
	num   int
	alive bool
}

const (
	playerNum = 100
	deadLine  = 4
)

func main() {
	r := ring.New(playerNum)
	for i := 0; i < playerNum; i++ {
		r.Value = &gamePlayer{
			num:   i + 1,
			alive: true,
		}
		r = r.Next()
	}

	callNum := 1 //报数
	deadPlayerNum := 0
	for deadPlayerNum < playerNum {
		r = r.Next()
		if r.Value.(*gamePlayer).alive {
			callNum++
		}
		if callNum == deadLine {
			r.Value.(*gamePlayer).alive = false
			fmt.Println("gamePlayer ", r.Value.(*gamePlayer).num, " is out")
			deadPlayerNum++
			callNum = 0
		}
	}
}
