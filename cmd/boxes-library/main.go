package main

import (
	"image/color"

	"github.com/mohanson/aiball-creator"
)

func mainShowLogo(aibo *aiball.AIBall) {
	aibo.Text(4, 6, "BOXES", 2)
	aibo.Join(200)
}

func mainRankInit(aibo *aiball.AIBall, rank [][]uint8) {
	for i := 0; i < aibo.Cols; i++ {
		for j := 0; j < aibo.Rows; j++ {
			aibo.Cell[j][i] = 4
		}
		aibo.Join(4)
	}
	for i := 0; i < aibo.Cols; i++ {
		for j := 0; j < aibo.Rows; j++ {
			aibo.Cell[j][i] = rank[j][i]
		}
		aibo.Join(4)
	}
}

func mainRank(aibo *aiball.AIBall, user []int, path string) {
	dest := make([]int, 2)
	back := make([]int, 2)
	for _, direct := range path {
		switch direct {
		case 'l':
			dest = []int{user[0], user[1] - 1}
			back = []int{user[0], user[1] - 2}
		case 'u':
			dest = []int{user[0] - 1, user[1]}
			back = []int{user[0] - 2, user[1]}
		case 'r':
			dest = []int{user[0], user[1] + 1}
			back = []int{user[0], user[1] + 2}
		case 'd':
			dest = []int{user[0] + 1, user[1]}
			back = []int{user[0] + 2, user[1]}
		}

		userstat := aibo.Cell[user[0]][user[1]]
		deststat := aibo.Cell[dest[0]][dest[1]]
		backstat := aibo.Cell[back[0]][back[1]]
		if userstat == 2 && deststat == 0 {
			aibo.Cell[user[0]][user[1]] = 0
			aibo.Cell[dest[0]][dest[1]] = 2
			copy(user, dest)
		}
		if userstat == 6 && deststat == 0 {
			aibo.Cell[user[0]][user[1]] = 4
			aibo.Cell[dest[0]][dest[1]] = 2
			copy(user, dest)
		}
		if userstat == 2 && deststat == 4 {
			aibo.Cell[user[0]][user[1]] = 0
			aibo.Cell[dest[0]][dest[1]] = 6
			copy(user, dest)
		}
		if userstat == 6 && deststat == 4 {
			aibo.Cell[user[0]][user[1]] = 4
			aibo.Cell[dest[0]][dest[1]] = 6
			copy(user, dest)
		}
		if userstat == 2 && deststat == 3 && backstat == 0 {
			aibo.Cell[user[0]][user[1]] = 0
			aibo.Cell[dest[0]][dest[1]] = 2
			aibo.Cell[back[0]][back[1]] = 3
			copy(user, dest)
		}
		if userstat == 6 && deststat == 3 && backstat == 0 {
			aibo.Cell[user[0]][user[1]] = 4
			aibo.Cell[dest[0]][dest[1]] = 2
			aibo.Cell[back[0]][back[1]] = 3
			copy(user, dest)
		}
		if userstat == 2 && deststat == 3 && backstat == 4 {
			aibo.Cell[user[0]][user[1]] = 0
			aibo.Cell[dest[0]][dest[1]] = 2
			aibo.Cell[back[0]][back[1]] = 5
			copy(user, dest)
		}
		if userstat == 6 && deststat == 3 && backstat == 4 {
			aibo.Cell[user[0]][user[1]] = 4
			aibo.Cell[dest[0]][dest[1]] = 2
			aibo.Cell[back[0]][back[1]] = 5
			copy(user, dest)
		}
		if userstat == 2 && deststat == 5 && backstat == 0 {
			aibo.Cell[user[0]][user[1]] = 0
			aibo.Cell[dest[0]][dest[1]] = 6
			aibo.Cell[back[0]][back[1]] = 3
			copy(user, dest)
		}
		if userstat == 6 && deststat == 5 && backstat == 0 {
			aibo.Cell[user[0]][user[1]] = 4
			aibo.Cell[dest[0]][dest[1]] = 6
			aibo.Cell[back[0]][back[1]] = 3
			copy(user, dest)
		}
		if userstat == 2 && deststat == 5 && backstat == 4 {
			aibo.Cell[user[0]][user[1]] = 0
			aibo.Cell[dest[0]][dest[1]] = 6
			aibo.Cell[back[0]][back[1]] = 5
			copy(user, dest)
		}
		if userstat == 6 && deststat == 5 && backstat == 4 {
			aibo.Cell[user[0]][user[1]] = 4
			aibo.Cell[dest[0]][dest[1]] = 6
			aibo.Cell[back[0]][back[1]] = 5
			copy(user, dest)
		}
		aibo.Join(16)
	}
	aibo.Join(84)
}

func mainWink(aibo *aiball.AIBall) {
	for i := 0; i < aibo.Rows/2; i++ {
		for j := 0; j < aibo.Cols; j++ {
			aibo.Cell[i][j] = 5
			aibo.Cell[aibo.Rows-1-i][j] = 5
		}
		aibo.Join(4)
	}
	aibo.Join(96)
}

func main() {
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
	aibo.Save("res/boxes-library.gif")
}
