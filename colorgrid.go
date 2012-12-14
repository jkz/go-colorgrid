package colorgrid

import "fmt"

type Color int
type ColorCode string

type Size struct {
	Width, Height int
}

type Grid struct {
	Cell Size
}

const (
	WHITE Color = iota
	BLACK
	RED
	GREEN
	LIGHT_BLUE
	MAGENTA
	YELLOW
)

var COLORS = map[Color]ColorCode{
	WHITE:      "\033[01;37m",
	BLACK:      "\033[22;30m",
	RED:        "\033[22;31m",
	GREEN:      "\033[22;32m",
	LIGHT_BLUE: "\033[01;34m",
	MAGENTA:    "\033[22;35m",
	YELLOW:     "\033[01;33m",
}

var BACKGROUND_COLORS = map[Color]ColorCode{
	WHITE:      "\033[01;47m",
	BLACK:      "\033[22;40m",
	RED:        "\033[22;41m",
	GREEN:      "\033[22;42m",
	LIGHT_BLUE: "\033[01;44m",
	MAGENTA:    "\033[22;45m",
	YELLOW:     "\033[01;43m",
}

func escape(c rune, val ...int) {
	fmt.Print("\033[", val[0])
	for _, d := range val[1:] {
		fmt.Printf(";", d)
	}
	fmt.Printf("%c", c)
	fmt.Println()
}

func jump(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}

func move(x, y int) {
	switch {
	case y < 0:
		fmt.Println("UP")
		fmt.Println("033[%dA", -y)
	case y > 0:
		fmt.Println("DOWN")
		fmt.Println("033[%dB", y)
	}
	switch {
	case x < 0:
		fmt.Println("RIGHT")
		fmt.Println("033[%dC", -x)
	case x > 0:
		fmt.Println("LEFT")
		fmt.Println("033[%dD", x)
	}
}

func Clear() {
	escape('J', 2)
}

func (g Grid) Render(x, y int, s string, c, bg Color) {
	fmt.Printf(string(COLORS[c]))
	fmt.Printf(string(BACKGROUND_COLORS[bg]))
	var xx, yy int
	for yy = 0; yy < g.Cell.Height; yy++ {
		for xx = 0; xx < g.Cell.Width; xx++ {
			jump(1+xx+x*g.Cell.Width, 1+yy+y*g.Cell.Height)
			if yy == g.Cell.Height/2 && xx == g.Cell.Width/2 {
				fmt.Print(s)
			} else {
				fmt.Print(" ")
			}
		}
	}
}
