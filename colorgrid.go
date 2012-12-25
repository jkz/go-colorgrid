package colorgrid

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type Color termbox.Attribute

type Size struct {
	Width, Height int
}

type Grid struct {
	Size   Size
	Fg, Bg Color
}

func (g Grid) fg() termbox.Attribute {
	return termbox.Attribute(g.Fg)
}

func (g Grid) bg() termbox.Attribute {
	return termbox.Attribute(g.Bg)
}

const (
	DEFAULT Color = iota
	BLACK
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

func (g Grid) Print(x, y int, s string, fg, bg Color) {
	termbox.SetCell(x, y, ' ', g.fg(), g.bg())
	termbox.SetCursor(x, y)
	fmt.Print(s)
}

func (g Grid) Cell(x, y int, ch rune, fg, bg Color) {
	var xx, yy, x_, y_ int
	var ch_ rune
	var fg_, bg_ termbox.Attribute
	for yy = 0; yy < g.Size.Height; yy++ {
		for xx = 0; xx < g.Size.Width; xx++ {
			if yy == g.Size.Height/2 && xx == g.Size.Width/2 {
				ch_ = ch
			} else {
				ch_ = ' '
			}
			x_ = 1 + xx + x*g.Size.Width
			y_ = 1 + yy + y*g.Size.Height
			fg_ = termbox.Attribute(fg)
			bg_ = termbox.Attribute(bg)
			//fmt.Printf("(%d, %d), %c, %d, %d\r\n", x_, y_, ch_, fg_, bg_)
			termbox.SetCell(x_, y_, ch_, fg_, bg_)
		}
	}
	termbox.SetCell(-1, -1, 'x', g.fg(), g.bg())
	termbox.Flush()
}

func (g Grid) Clear() {
	termbox.Clear(g.fg(), g.bg())
}

func (g Grid) Flush() {
	termbox.Flush()
}

func NewGrid(width, height int, fg, bg Color) Grid {
	return Grid{Size{width, height}, fg, bg}
}
