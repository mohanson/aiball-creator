package main

import (
	"image/color"
	"image/gif"
	"os"

	"github.com/godump/doa"
	"github.com/mohanson/aiball-creator"
)

func main() {
	cell := make([][]int, 45)
	for i := 0; i < 45; i++ {
		cell[i] = make([]int, 80)
	}
	aibo := aiball.AIBall{
		Rows:     27,
		Cols:     48,
		Cell:     cell,
		RGBABase: color.RGBA{0xff, 0xff, 0xff, 0xff},
		RGBA: []color.RGBA{
			{0xeb, 0xed, 0xf0, 0xff},
			{0xd1, 0xba, 0x74, 0xff},
			{0xf4, 0x60, 0x6c, 0xff},
			{0x19, 0xca, 0xad, 0xff},
			{0x9c, 0x9c, 0x9c, 0xff},
			{0x2f, 0x4f, 0x4f, 0xff},
			{0xf4, 0x60, 0x6c, 0xff},
		},
		CellSize: 18,
		CellPads: 2,
	}
	aibo.Cell[4][2] = 2
	pale := aibo.Draw()

	file := doa.Try(os.OpenFile("/tmp/foo.gif", os.O_WRONLY|os.O_TRUNC, 0755))
	defer file.Close()
	anim := gif.GIF{}
	anim.Delay = append(anim.Delay, 8)
	anim.Image = append(anim.Image, pale)
	gif.EncodeAll(file, &anim)
}
