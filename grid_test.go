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

func TestAdjacentCoordinatesSouthIsDefined(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	got := grid.AdjacentCoordinates(
		mazegen.South,
		&mazegen.Coordinates{
			Y: 0,
			X: 0,
		},
	)

	if got == nil {
		t.Error("response is nil which is unexpected")
		return
	}

	if got.Y != 1 {
		t.Errorf("response has unexpected Y value=%d", got.Y)
	}

	if got.X != 0 {
		t.Errorf("response has unexpected X value=%d", got.X)
	}
}

func TestAdjacentCoordinatesEastIsDefined(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	got := grid.AdjacentCoordinates(
		mazegen.East,
		&mazegen.Coordinates{
			Y: 0,
			X: 0,
		},
	)

	if got == nil {
		t.Error("response is nil which is unexpected")
		return
	}

	if got.Y != 0 {
		t.Errorf("response has unexpected Y value=%d", got.Y)
	}

	if got.X != 1 {
		t.Errorf("response has unexpected X value=%d", got.X)
	}
}

func TestAdjacentCoordinatesWestIsNil(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	got := grid.AdjacentCoordinates(
		mazegen.West,
		&mazegen.Coordinates{
			Y: 0,
			X: 0,
		},
	)

	if got != nil {
		t.Error("response is not nil which is unexpected")
	}

}

func TestAdjacentCellSouthIsDefined(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	got := grid.AdjacentCell(
		mazegen.South,
		&mazegen.Coordinates{
			Y: 0,
			X: 0,
		},
	)

	if got == nil {
		t.Error("response is nil which is unexpected")
		return
	}
}

func TestAdjacentCellEastIsDefined(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	got := grid.AdjacentCell(
		mazegen.East,
		&mazegen.Coordinates{
			Y: 0,
			X: 0,
		},
	)

	if got == nil {
		t.Error("response is nil which is unexpected")
		return
	}
}

func TestAdjacentCellWestIsNil(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	got := grid.AdjacentCell(
		mazegen.West,
		&mazegen.Coordinates{
			Y: 0,
			X: 0,
		},
	)

	if got != nil {
		t.Error("response is not nil which is unexpected")
		return
	}
}

func TestCoordinatesInBoundsReturnsTrue(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	got := grid.CoordinatesInBounds(
		&mazegen.Coordinates{
			Y: 1,
			X: 1,
		},
	)

	if got == false {
		t.Error("response is false but should be true")
	}
}

func TestCoordinatesInBoundsReturnsFalse(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	got := grid.CoordinatesInBounds(
		&mazegen.Coordinates{
			Y: 2,
			X: 2,
		},
	)

	if got == true {
		t.Error("response is true but should be false")
	}
}

func TestRowInBoundsReturnsFalse(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	responses := []bool{
		grid.RowInBounds(-1),
		grid.RowInBounds(2),
	}

	for index, value := range responses {
		if value == true {
			t.Errorf("response is unexpected index=%d", index)
		}
	}
}

func TestRowInBoundsReturnsTrue(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	responses := []bool{
		grid.RowInBounds(0),
		grid.RowInBounds(1),
	}

	for index, value := range responses {
		if value == false {
			t.Errorf("response is unexpected index=%d", index)
		}
	}
}

func TestColumnInBoundsReturnsFalse(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	responses := []bool{
		grid.ColumnInBounds(-1),
		grid.ColumnInBounds(2),
	}

	for index, value := range responses {
		if value == true {
			t.Errorf("response is unexpected index=%d", index)
		}
	}
}

func TestColumnInBoundsReturnsTrue(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	responses := []bool{
		grid.ColumnInBounds(0),
		grid.ColumnInBounds(1),
	}

	for index, value := range responses {
		if value == false {
			t.Errorf("response is unexpected index=%d", index)
		}
	}
}
