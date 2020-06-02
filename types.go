package main

import "fmt"

type Cell struct {
	X       int
	Y       int
	HasMine bool
}

type Board struct {
	Width  int
	Height int
	Cells  map[string]*Cell
}

func (b *Board) GetCellAtLinearIndex(index int) *Cell {
	x, y := getCoordinatesOfLinearIndex(index, b.Width)
	k := getKey(x, y)
	return b.Cells[k]
}

func getCoordinatesOfLinearIndex(index, width int) (int, int) {
	y := index / width
	x := index - (width*y)
	return x, y
}

func NewBoard(width, height, mineCount int) *Board {
	cells := map[string]*Cell{}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := &Cell{X: x, Y: y, HasMine: false}
			k := getKey(x, y)
			cells[k] = c
		}
	}
	return &Board{
		Width:  width,
		Height: height,
		Cells:  cells,
	}
}

func getKey(x, y int) string {
	return fmt.Sprintf("%d_%d", x, y)
}