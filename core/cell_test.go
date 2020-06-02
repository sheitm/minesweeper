package core

import (
	"testing"
)

func TestCell_getNeighbors(t *testing.T) {
	// Arrange
	b := NewBoard(6, 4, 6)
	c := b.GetCellAt(1, 1)

	expectedCoords := []string{
		"0_0",
		"1_0",
		"2_0",
		"0_1",
		"2_1",
		"0_2",
		"1_2",
		"2_2",
	}

	// Act
	n := c.getNeighbors()

	// Assert
	if len(n) != 8  {
		t.Errorf("unexpected number of neighbors, got %d", len(n))
	}

	var coords []string
	for _, cell := range n {
		k := getKey(cell.X, cell.Y)
		coords = append(coords, k)
	}
	for _, coord := range expectedCoords {
		if !checkExists(coord, coords) {
			t.Errorf("didn't fin coordinates %s", coord)
		}
	}
}

func checkExists(coord string, coords[]string) bool {
	for _, c := range coords {
		if c == coord {
			return true
		}
	}
	return false
}