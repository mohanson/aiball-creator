package main

import (
	"flag"
	"image/color"

	"github.com/mohanson/aiball-creator"
)

func mainP0() {
	aibo := aiball.NewAIBall(64, 36)
	aibo.RGBABase = color.RGBA{0xff, 0xff, 0xff, 0xff}
	aibo.RGBA = []color.RGBA{
		{0xeb, 0xed, 0xf0, 0xff},
		{0xd1, 0xba, 0x74, 0xff},
		{0xf4, 0x60, 0x6c, 0xff},
		{0x19, 0xca, 0xad, 0xff},
		{0x9c, 0x9c, 0x9c, 0xff},
		{0x2f, 0x4f, 0x4f, 0xff},
		{0xf4, 0x60, 0x6c, 0xff},
	}
	aibo.CellSize = 13
	aibo.CellPads = 2

	data := [][]uint8{
		{1, 1, 1, 1, 1, 0, 0, 0, 0},
		{1, 2, 0, 0, 1, 0, 0, 0, 0},
		{1, 0, 3, 3, 1, 0, 1, 1, 1},
		{1, 0, 3, 0, 1, 0, 1, 4, 1},
		{1, 1, 1, 0, 1, 1, 1, 4, 1},
		{0, 1, 1, 0, 0, 0, 0, 4, 1},
		{0, 1, 0, 0, 0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 0, 0, 0},
	}
	conv := [][]uint8{}
	for i := 0; i < len(data)*3; i++ {
		conv = append(conv, make([]uint8, len(data[0])*3))
	}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			for k := 0; k < 2; k++ {
				for l := 0; l < 2; l++ {
					conv[i*2+k][j*2+l] = data[i][j]
				}
			}
		}
	}
	aibo.Copy(23, 4, conv)
	aibo.Text(20, 26, "BOXES", 2)
	aibo.Join(0)
	aibo.Save("res/boxes-p0.png")
}

func mainP1() {
	aibo := aiball.NewAIBall(40, 60)
	aibo.RGBABase = color.RGBA{0xff, 0xff, 0xff, 0xff}
	aibo.RGBA = []color.RGBA{
		{0xeb, 0xed, 0xf0, 0xff},
		{0xd1, 0xba, 0x74, 0xff},
		{0xf4, 0x60, 0x6c, 0xff},
		{0x19, 0xca, 0xad, 0xff},
		{0x9c, 0x9c, 0x9c, 0xff},
		{0x2f, 0x4f, 0x4f, 0xff},
		{0xf4, 0x60, 0x6c, 0xff},
	}
	aibo.CellSize = 13
	aibo.CellPads = 2

	data := [][]uint8{
		{1, 1, 1, 1, 1, 0, 0, 0, 0},
		{1, 2, 0, 0, 1, 0, 0, 0, 0},
		{1, 0, 3, 3, 1, 0, 1, 1, 1},
		{1, 0, 3, 0, 1, 0, 1, 4, 1},
		{1, 1, 1, 0, 1, 1, 1, 4, 1},
		{0, 1, 1, 0, 0, 0, 0, 4, 1},
		{0, 1, 0, 0, 0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 0, 0, 0},
	}
	conv := [][]uint8{}
	for i := 0; i < len(data)*3; i++ {
		conv = append(conv, make([]uint8, len(data[0])*3))
	}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					conv[i*3+k][j*3+l] = data[i][j]
				}
			}
		}
	}
	aibo.Copy(7, 10, conv)
	aibo.Text(8, 45, "BOXES", 2)
	aibo.Join(0)
	aibo.Save("res/boxes-p1.png")
}

func main() {
	flag.Parse()
	switch flag.Arg(0) {
	case "p0":
		mainP0()
	case "p1":
		mainP1()
	}
}
