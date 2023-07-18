package aiball

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
	"os/exec"

	"github.com/godump/doa"
)

type AIBall struct {
	Rows     int
	Cols     int
	Cell     [][]uint8
	RGBABase color.RGBA
	RGBA     []color.RGBA
	CellSize int
	CellPads int
	GIFs     gif.GIF
}

func NewAIBall(x int, y int) *AIBall {
	cell := make([][]uint8, y)
	for i := 0; i < y; i++ {
		cell[i] = make([]uint8, x)
	}
	gifs := gif.GIF{}
	return &AIBall{
		Rows: y,
		Cols: x,
		Cell: cell,
		GIFs: gifs,
	}
}

func (a *AIBall) Text(x int, y int, data string, cell uint8) {
	for _, c := range data {
		char := Font[c]
		w := len(char[0])
		for i := 0; i < len(char); i++ {
			for j := 0; j < w; j++ {
				if char[i][j] != 0 {
					a.Cell[y+i][x+j] = cell
				}
			}
		}
		x += w + 1
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

func (a *AIBall) Join(d int) {
	a.GIFs.Image = append(a.GIFs.Image, a.Draw())
	a.GIFs.Delay = append(a.GIFs.Delay, d)
}

func (a *AIBall) Save(name string) {
	file := doa.Try(os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755))
	defer file.Close()
	gif.EncodeAll(file, &a.GIFs)
	if _, err := exec.LookPath("gifsicle"); err == nil {
		doa.Nil(exec.Command("gifsicle", "-O3", name, "-o", name).Run())
	}
}
