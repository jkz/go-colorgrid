package main

import "github.com/jessethegame/colorgrid"

func main() {
	g := colorgrid.Grid{colorgrid.Size{5, 3}, colorgrid.RED, colorgrid.GREEN}
	g.Clear()
	var x, y int
	for y = 0; y < 8; y++ {
		for x = 0; x < 8; x++ {
			if (x%2 == 0 && y%2 == 0) || (x%2 == 1 && y%2 == 1) {
				g.Print(x, y, "X", colorgrid.BLACK, colorgrid.WHITE)
			} else {
				g.Print(x, y, "X", colorgrid.WHITE, colorgrid.BLACK)
			}
		}
	}
}
