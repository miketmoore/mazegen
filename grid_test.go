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

func TestCoordinatesInBounds(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	tests := []struct {
		x, y     int
		expected bool
	}{
		{x: -1, y: -1, expected: false},
		{x: 0, y: 0, expected: true},
		{x: 1, y: 0, expected: true},
		{x: 1, y: 1, expected: true},
		{x: 0, y: 1, expected: true},
		{x: 2, y: 2, expected: false},
	}

	for index, test := range tests {
		got := grid.CoordinatesInBounds(&mazegen.Coordinates{
			X: test.x,
			Y: test.y,
		})
		if got != test.expected {
			t.Errorf("test failed index=%d", index)
		}
	}
}

func TestRowInBounds(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	tests := []struct {
		input    int
		expected bool
	}{
		{input: -1, expected: false},
		{input: 0, expected: true},
		{input: 1, expected: true},
		{input: 2, expected: false},
	}

	for index, test := range tests {
		got := grid.RowInBounds(test.input)
		if got != test.expected {
			t.Errorf("test failed index=%d", index)
		}
	}
}

func TestColumnInBounds(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	tests := []struct {
		input    int
		expected bool
	}{
		{input: -1, expected: false},
		{input: 0, expected: true},
		{input: 1, expected: true},
		{input: 2, expected: false},
	}

	for index, test := range tests {
		got := grid.ColumnInBounds(test.input)
		if got != test.expected {
			t.Errorf("test failed index=%d", index)
		}
	}
}

func TestRandomCoordinates(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	for i := 0; i < 100; i++ {
		got := grid.RandomCoordinates()
		if got == nil {
			t.Errorf("response is nil which is unexpected index=%d", i)
		}
	}
}

func TestRandomCell(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	for i := 0; i < 100; i++ {
		got := grid.RandomCell()
		if got == nil {
			t.Errorf("response is nil which is unexpected index=%d", i)
		}
	}
}

func TestIsWallAvailable(t *testing.T) {

	tests := []struct {
		grid        *mazegen.Grid
		cell        *mazegen.Cell
		coordinates *mazegen.Coordinates
		direction   mazegen.DirectionValue
		expected    bool
	}{
		// top left
		{
			grid:        buildGrid(t, 2, 2),
			cell:        mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{Y: 0, X: 0},
			direction:   mazegen.North,
			expected:    false,
		},
		{
			grid:        buildGrid(t, 2, 2),
			cell:        mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{Y: 0, X: 0},
			direction:   mazegen.East,
			expected:    true,
		},
		{
			grid:        buildGrid(t, 2, 2),
			cell:        mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{Y: 0, X: 0},
			direction:   mazegen.South,
			expected:    true,
		},
		{
			grid:        buildGrid(t, 2, 2),
			cell:        mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{Y: 0, X: 0},
			direction:   mazegen.West,
			expected:    false,
		},

		// top right
		{
			grid:        buildGrid(t, 2, 2),
			cell:        mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{Y: 0, X: 1},
			direction:   mazegen.North,
			expected:    false,
		},
		{
			grid:        buildGrid(t, 2, 2),
			cell:        mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{Y: 0, X: 1},
			direction:   mazegen.East,
			expected:    false,
		},
		{
			grid:        buildGrid(t, 2, 2),
			cell:        mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{Y: 0, X: 1},
			direction:   mazegen.South,
			expected:    true,
		},
		{
			grid:        buildGrid(t, 2, 2),
			cell:        mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{Y: 0, X: 1},
			direction:   mazegen.West,
			expected:    true,
		},

		// bottom left
		{
			grid:        buildGrid(t, 2, 2),
			cell:        mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{Y: 1, X: 0},
			direction:   mazegen.North,
			expected:    true,
		},
		{
			grid:        buildGrid(t, 2, 2),
			cell:        mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{Y: 1, X: 0},
			direction:   mazegen.East,
			expected:    true,
		},
		{
			grid:        buildGrid(t, 2, 2),
			cell:        mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{Y: 1, X: 0},
			direction:   mazegen.South,
			expected:    false,
		},
		{
			grid:        buildGrid(t, 2, 2),
			cell:        mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{Y: 1, X: 0},
			direction:   mazegen.West,
			expected:    false,
		},

		// bottom right
		{
			grid:        buildGrid(t, 2, 2),
			cell:        mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{Y: 1, X: 1},
			direction:   mazegen.North,
			expected:    true,
		},
		{
			grid:        buildGrid(t, 2, 2),
			cell:        mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{Y: 1, X: 1},
			direction:   mazegen.East,
			expected:    false,
		},
		{
			grid:        buildGrid(t, 2, 2),
			cell:        mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{Y: 1, X: 1},
			direction:   mazegen.West,
			expected:    true,
		},
		{
			grid:        buildGrid(t, 2, 2),
			cell:        mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{Y: 1, X: 1},
			direction:   mazegen.South,
			expected:    false,
		},
	}

	for index, test := range tests {
		got := test.grid.IsWallAvailable(
			test.coordinates,
			test.direction,
			test.cell,
		)
		if got != test.expected {
			t.Errorf("test failed index=%d", index)
		}
	}
}

func TestAvailableCellWalls(t *testing.T) {

	tests := []struct {
		grid        *mazegen.Grid
		cell        *mazegen.Cell
		coordinates *mazegen.Coordinates
		expected    []mazegen.DirectionValue
	}{
		{
			grid: buildGrid(t, 2, 2),
			cell: mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{
				Y: 0,
				X: 0,
			},
			expected: []mazegen.DirectionValue{
				mazegen.East,
				mazegen.South,
			},
		},
		{
			grid: buildGrid(t, 2, 2),
			cell: mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{
				Y: 1,
				X: 0,
			},
			expected: []mazegen.DirectionValue{
				mazegen.North,
				mazegen.East,
			},
		},
		{
			grid: buildGrid(t, 2, 2),
			cell: mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{
				Y: 0,
				X: 1,
			},
			expected: []mazegen.DirectionValue{
				mazegen.South,
				mazegen.West,
			},
		},
		{
			grid: buildGrid(t, 2, 2),
			cell: mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{
				Y: 1,
				X: 1,
			},
			expected: []mazegen.DirectionValue{
				mazegen.North,
				mazegen.West,
			},
		},
		{
			grid: buildGrid(t, 3, 3),
			cell: mazegen.NewCell(),
			coordinates: &mazegen.Coordinates{
				Y: 1,
				X: 1,
			},
			expected: []mazegen.DirectionValue{
				mazegen.North,
				mazegen.East,
				mazegen.South,
				mazegen.West,
			},
		},
	}

	for index, test := range tests {
		got := test.grid.AvailableCellWalls(
			test.cell,
			test.coordinates,
		)
		if len(got) != len(test.expected) {
			t.Errorf("test failed index=%d", index)
		}
		for i := 0; i < len(test.expected); i++ {
			if got[i] != test.expected[i] {
				t.Errorf("test failed index=%d", index)
			}
		}
	}
}

func TestUpdateCell(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	coordinates := mazegen.NewCoordinates(0, 0)

	cell := mazegen.NewCell()
	cell.Visited = true

	grid.SetCell(
		coordinates,
		cell,
	)

	cell = grid.Cell(coordinates)
	if cell.Visited != true {
		t.Error("test failed")
	}
}

func TestCarveCellWall(t *testing.T) {
	grid := buildGrid(t, 2, 2)

	coordinates := mazegen.NewCoordinates(0, 0)

	err := grid.CarveCellWall(
		coordinates,
		mazegen.North,
	)

	if err != nil {
		t.Error("error is unexpected")
	}

	cell := grid.Cell(coordinates)
	if cell.IsWallSolid(mazegen.North) {
		t.Error("wall should not be solid")
	}
}
