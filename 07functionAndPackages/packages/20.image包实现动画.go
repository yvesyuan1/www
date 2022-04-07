package main

/*
	这个go文件直接在golang里运行无法运行
	正确做法：去终端编译该go文件，获得Unix可执行文件，
	在终端运行该可执行文件，并生成.gif文件
	//	go build 20.image包实现动画.go
	//	./20.image包实现动画 >1.gif
*/

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black} //调色板
const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) //随机数种子
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		cycles  = 5     //完整的x振荡器变化的个数
		res     = 0.001 //角度分辨率
		size    = 100   //图像画布包含[-size. .+size]
		nframes = 64    //动画中的帧数
		delay   = 8     //以10ms为单位的帧间延迟
	)
	freq := rand.Float64() * 3.0        //y 振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes} //创建gif
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) //创建区域
		img := image.NewPaletted(rect, palette)      //创建图片
		for t := 0.0; t < cycles*8*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
