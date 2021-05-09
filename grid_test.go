package mazegen_test

import (
	"fmt"
	"testing"

	"github.com/miketmoore/mazegen"
)

func buildGrid(t *testing.T, rows, columns int) *mazegen.Grid {
	grid, err := mazegen.NewGrid(rows, columns, mazegen.NewRandom())
	if err != nil {
		fmt.Println(err)
		t.Error("creating grid failed")
		return nil
	}
	return grid
}

func TestNewGrid(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	if grid.Rows != 2 {
		t.Error("unexpected value for grid.Rows")
	}

	if grid.Columns != 2 {
		t.Error("unexpected value for grid.Columns")
	}

	if grid.Random == nil {
		t.Error("grid.Random is nil which is unexpected")
	}

	if len(grid.Cells) != 2 {
		t.Error("unexpected total rows")
	}

	if len(grid.Cells[0]) != 2 {
		t.Error("unexpected total cells in first row")
	}

	if len(grid.Cells[1]) != 2 {
		t.Error("unexpected total cells in second row")
	}

	for rowIndex, rowSlice := range grid.Cells {
		for columnIndex, cell := range rowSlice {
			if cell == nil {
				t.Errorf("rowIndex=%d columnIndex=%d value is nil which is unexpected", rowIndex, columnIndex)
			}
		}
	}
}

func TestCell(t *testing.T) {
	grid := buildGrid(t, 2, 2)
	cell := grid.Cell(&mazegen.Coordinates{
		Y: 0,
		X: 0,
	})
	if cell == nil {
		t.Error("cell is nil which is unexpected")
	}

	cell = grid.Cell(&mazegen.Coordinates{
		Y: 0,
		X: 1,
	})
	if cell == nil {
		t.Error("cell is nil which is unexpected")
	}

	cell = grid.Cell(&mazegen.Coordinates{
		Y: 1,
		X: 0,
	})
	if cell == nil {
		t.Error("cell is nil which is unexpected")
	}

	cell = grid.Cell(&mazegen.Coordinates{
		Y: 1,
		X: 1,
	})
	if cell == nil {
		t.Error("cell is nil which is unexpected")
	}

	// Test getting a cell out of bounds
	// should return nil
	cell = grid.Cell(&mazegen.Coordinates{
		Y: 2,
		X: 2,
	})
	if cell != nil {
		t.Error("cell is not nil which is unexpected")
	}
}
