package core

type Cell struct {
	X                 int
	Y                 int
	HasMine           bool
	State             string // Closed, Opened, Failed
	NeighborMineCount int
	board             *Board
}

func (c *Cell) setMine() {
	c.HasMine = true
}

func (c *Cell) setNeighborMineCount() {
	count := 0
	for _, cell := range c.getNeighbors() {
		if cell.HasMine {
			count++
		}
	}
	c.NeighborMineCount = count
}

func (c *Cell) getNeighbors() []*Cell {
	var neighbors []*Cell
	yOffset := -1
	xOffset := -1
	for i := 0; i < 9; i++ {
		y := c.Y + yOffset
		x := c.X + xOffset
		n := c.board.GetCellAt(x, y)
		if n != nil && n != c {
			neighbors = append(neighbors, n)
		}
		xOffset++
		if xOffset > 1 {
			xOffset = -1
			yOffset++
		}
	}
	return neighbors
}