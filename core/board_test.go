package core

import "testing"


func TestBoard_GetCellAtLinearIndex(t *testing.T) {
	// Arrange
	b := NewBoard(6, 4, 0)

	// Act
	c := b.GetCellAtLinearIndex(0)
	if c.X != 0 {
		t.Errorf("unexpected x value, expected 0, got %d", c.X)
	}
	if c.Y != 0 {
		t.Errorf("unexpected x value, expected 0, got %d", c.Y)
	}

	c = b.GetCellAtLinearIndex(5)
	if c.X != 5 {
		t.Errorf("unexpected x value, expected 5, got %d", c.X)
	}
	if c.Y != 0 {
		t.Errorf("unexpected x value, expected 0, got %d", c.Y)
	}
}

func Test_getCoordinatesOfLinearIndex(t *testing.T) {
	type args struct {
		index  int
		width  int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"1", args{0, 6}, 0, 0},
		{"2", args{5, 6}, 5, 0},
		{"3", args{6, 6}, 0, 1},
		{"4", args{7, 6}, 1, 1},
		{"5", args{23, 6}, 5, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getCoordinatesOfLinearIndex(tt.args.index, tt.args.width)
			if got != tt.want {
				t.Errorf("getCoordinatesOfLinearIndex() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getCoordinatesOfLinearIndex() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}