package aiball

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
	"time"

	"github.com/godump/doa"
)

type AIBall struct {
	Rows     int
	Cols     int
	Cell     [][]int
	RGBABase color.RGBA
	RGBA     []color.RGBA
	CellSize int
	CellPads int
	GIFs     gif.GIF
}

func NewAIBall(r int, c int) *AIBall {
	cell := make([][]int, r)
	for i := 0; i < r; i++ {
		cell[i] = make([]int, c)
	}
	return &AIBall{
		Rows: r,
		Cols: c,
		Cell: cell,
		GIFs: gif.GIF{},
	}
}

func (a *AIBall) Draw() *image.Paletted {
	w := a.Cols * (a.CellSize + a.CellPads)
	h := a.Rows * (a.CellSize + a.CellPads)
	m := image.NewPaletted(image.Rect(0, 0, w, h), palette.Plan9)
	draw.Draw(m, m.Bounds(), &image.Uniform{a.RGBABase}, image.Point{}, draw.Src)
	for i := 0; i < a.Rows; i++ {
		for j := 0; j < a.Cols; j++ {
			x := a.CellPads/2 + j*(a.CellSize+a.CellPads)
			y := a.CellPads/2 + i*(a.CellSize+a.CellPads)
			c := image.Rect(x, y, x+a.CellSize, y+a.CellSize)
			draw.Draw(m, c.Bounds(), &image.Uniform{a.RGBA[a.Cell[i][j]]}, image.Point{x, y}, draw.Src)
		}
	}
	return m
}

func (a *AIBall) Join(d time.Duration) {
	a.GIFs.Image = append(a.GIFs.Image, a.Draw())
	a.GIFs.Delay = append(a.GIFs.Delay, int(d.Milliseconds())/10)
}

func (a *AIBall) Save(name string) {
	file := doa.Try(os.OpenFile(name, os.O_WRONLY|os.O_TRUNC, 0755))
	defer file.Close()
	gif.EncodeAll(file, &a.GIFs)
}
