package main

import (
	"github.com/sheitm/minesweeper/core"
	"github.com/sheitm/minesweeper/ui"
)

func main() {
	b := core.NewBoard(6, 4, 6)
	ui.PrintBoard(b)
}