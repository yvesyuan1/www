package main

import (
	"github.com/skip2/go-qrcode"
)

func main() {
	qrcode.WriteFile("BossNiu", qrcode.Medium, 256, "./qrcode.png")

}
