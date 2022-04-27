package main

import (
	"container/ring"
	"fmt"
)

type Player struct {
	pos   int  //位置
	alive bool //是否或者
}

const (
	playerCount = 100
	startPos    = 1
	deadline    = 3 //数到3就dead
)

func main() {
	r := ring.New(playerCount)
	for i := startPos; i <= playerCount; i++ {
		r.Value = &Player{i, true}
		r = r.Next()
	}
	if startPos > 1 {
		r = r.Move(startPos - 1)
	}

	counter := 1                  //从1开始报数
	deadCount := 0                //死亡的人数
	for deadCount < playerCount { //没有死光，继续
		r = r.Next() //跳到下一个
		if r.Value.(*Player).alive {
			counter++
		}
		if counter == deadline { //3，这个人挂掉
			r.Value.(*Player).alive = false
			fmt.Printf("Player%d被扔进大海\n", r.Value.(*Player).pos)
			deadCount++ //死亡一个
			counter = 0 //报数清零
		}
	}

}
