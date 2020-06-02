package ui

import (
	"fmt"
	"github.com/sheitm/minesweeper/core"
)

func PrintBoard(board *core.Board) {
	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			c := board.GetCellAt(x, y)
			if c.HasMine {
				fmt.Print("* ")
			} else {
				fmt.Printf("%d ", c.NeighborMineCount)
			}
		}
		fmt.Println()
	}
}
