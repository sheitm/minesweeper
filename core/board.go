package core

import (
	"fmt"
	"math/rand"
	"time"
)

func NewBoard(width, height, mineCount int) *Board {
	rand.Seed(time.Now().UnixNano())

	b :=  &Board{
		Width:  width,
		Height: height,
	}

	cells := map[string]*Cell{}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := &Cell{X: x, Y: y, HasMine: false, board: b}
			k := getKey(x, y)
			cells[k] = c
		}
	}

	b.Cells = cells
	b.placeMines(mineCount)

	return b
}

type Board struct {
	Width  int
	Height int
	Cells  map[string]*Cell
}

func (b *Board) GetCellAtLinearIndex(index int) *Cell {
	x, y := getCoordinatesOfLinearIndex(index, b.Width)
	return b.GetCellAt(x, y)
}

func (b *Board) GetCellAt(x, y int) *Cell {
	if x < 0 || y < 0 || x >= b.Width || y >= b.Height {
		return nil
	}
	k := getKey(x, y)
	return b.Cells[k]
}

func (b *Board) placeMines(c int) {
	for i := 0; i < c; i++ {
		b.placeMine()
	}
	for i := 0; i < b.Height*b.Width; i++ {
		b.GetCellAtLinearIndex(i).setNeighborMineCount()
	}
}

func (b *Board) placeMine() {
	for {
		index := rand.Intn(b.Width * b.Height)
		c := b.GetCellAtLinearIndex(index)
		if !c.HasMine {
			c.setMine()
			return
		}
	}
}

func getCoordinatesOfLinearIndex(index, width int) (int, int) {
	y := index / width
	x := index - (width*y)
	return x, y
}

func getKey(x, y int) string {
	return fmt.Sprintf("%d_%d", x, y)
}