package main

import (
	"github.com/skip2/go-qrcode"
	"image/color"
	"log"
)

func main() {
	qr, err := qrcode.New("BossNiu", qrcode.Medium) //这里qr是二维码的图片字节
	if err != nil {
		log.Fatal(err)
	} else {
		qr.BackgroundColor = color.RGBA{50, 205, 50, 255}
		qr.ForegroundColor = color.White
		qr.WriteFile(256, "./06functionAndPackages/packages/golang_qrcode.png")
	}

}
