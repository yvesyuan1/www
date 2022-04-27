package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	const size = 300                                   //图像大小
	pic := image.NewGray(image.Rect(0, 0, size, size)) //根据大小创建一个灰度图
	//遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			//填充白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	//从0到最大像素生成x坐标
	for x := 0; x < size; x++ {
		//让sin的值的范围在0 - 2Pi之间
		s := float64(x) * 2 * math.Pi / size
		//sin的幅度为一般的像素，乡下便宜一半像素并旋转
		y := size/2 - math.Cos(s)*size/2
		//用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}
	file, _ := os.Create("cos.png")
	png.Encode(file, pic) //写入图像
	file.Close()
}
