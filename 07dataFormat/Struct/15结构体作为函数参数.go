package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Subjects struct {
	chinese float64
	math    float64
	english float64
}

type Student15 struct {
	id       int
	name     string
	subjects Subjects
	avgScore float64
}

func getSubjSum(subjects Subjects) float64 { //值传递
	return subjects.chinese + subjects.math + subjects.english
}

func getStuAvg(p *Student15) { //地址传递
	subjSum := getSubjSum(p.subjects)
	p.avgScore = subjSum / 3
}

func main() {
	students := [5]Student15{}
	for i := 0; i < 5; i++ {
		rand.Seed(int64(i))
		students[i].id = i
		students[i].name = fmt.Sprintf("学生%d", i+1)
		//fmt.Sprintf("%.2f",)先将float64转为对应格式的string，再用strconv.ParseFloat(,64）转回float64
		students[i].subjects.chinese, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float64()*100), 64)
		students[i].subjects.math, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float64()*100), 64)
		students[i].subjects.english, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float64()*100), 64)
		getStuAvg(&students[i])
		fmt.Println(students[i])
	}
}
