package mazegen

import (
	"fmt"
	"math/rand"
	"time"
)

type Grid struct {
	Rows, Columns   int
	Cells           [][]string
	CellDataManager *CellDataManager
	Random          *rand.Rand
}

func NewRandom() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func NewGrid(rows, columns int, random *rand.Rand) (*Grid, error) {
	grid := &Grid{
		Rows:            rows,
		Columns:         columns,
		Cells:           [][]string{},
		CellDataManager: NewCellDataManager(),
		Random:          random,
	}

	for row := 0; row < rows; row++ {
		grid.Cells[row] = []string{}
		for column := 0; column < columns; column++ {
			cell := NewCell()
			cellData, err := grid.CellDataManager.Data(cell)
			if err != nil {
				fmt.Println(err)
				return nil, fmt.Errorf("error getting cell data")
			}
			grid.Cells[row][column] = cellData
		}
	}

	return grid, nil
}

func (grid *Grid) ForEachRow(
	callback func(row []*Cell, rowIndex int),
) {
	for rowIndex, row := range grid.Cells {
		rowToReturn := []*Cell{}
		for _, cellData := range row {
			cell := grid.CellDataManager.NewFromData(cellData)
			rowToReturn = append(rowToReturn, cell)
		}
		callback(rowToReturn, rowIndex)
	}
}

func (grid *Grid) Cell(coordinates *Coordinates) *Cell {
	row := coordinates.Y
	if row >= 0 && row < len(grid.Cells) {
		cellsRow := grid.Cells[row]
		column := coordinates.X
		if column >= 0 && column < len(cellsRow) {
			data := cellsRow[column]
			return grid.CellDataManager.NewFromData(data)
		}
	}
	return nil
}

func (grid *Grid) AdjacentCoordinates(direction DirectionValue, coordinates *Coordinates) *Coordinates {
	row := coordinates.Y
	column := coordinates.X
	switch direction {
	case North:
		if row == 0 {
			return nil
		}
		return &Coordinates{Y: row - 1, X: column}
	case East:
		if column == grid.Columns-1 {
			return nil
		}
		return &Coordinates{Y: row, X: column + 1}
	case South:
		if row == grid.Rows-1 {
			return nil
		}
		return &Coordinates{Y: row + 1, X: column}
	case West:
		if column == 0 {
			return nil
		}
		return &Coordinates{Y: row, X: column - 1}
	}

	return nil
}

func (grid *Grid) AdjacentCell(direction DirectionValue, coordinates *Coordinates) *Cell {
	adjacentCoordinates := grid.AdjacentCoordinates(direction, coordinates)
	if adjacentCoordinates == nil {
		return nil
	}
	if grid.CoordinatesInBounds(adjacentCoordinates) {
		return grid.Cell(adjacentCoordinates)
	}
	return nil
}

func (grid *Grid) CoordinatesInBounds(coordinates *Coordinates) bool {
	return grid.RowInBounds(coordinates.Y) && grid.ColumnInBounds(coordinates.X)
}

func (grid *Grid) RowInBounds(row int) bool {
	return row >= 0 && row < grid.Rows
}

func (grid *Grid) ColumnInBounds(column int) bool {
	return column >= 0 && column < grid.Columns
}

func (grid *Grid) RandomCoordinates() *Coordinates {
	// TODO
	x := grid.Random.Intn(grid.Rows - 1)
	y := grid.Random.Intn(grid.Columns - 1)
	return &Coordinates{
		X: x,
		Y: y,
	}
}

func (grid *Grid) RandomCell() *Cell {
	coordinates := grid.RandomCoordinates()
	cellData := grid.Cells[coordinates.Y][coordinates.X]
	return grid.CellDataManager.NewFromData(cellData)
}

func (grid *Grid) IsWallAvailable(
	coordinates *Coordinates,
	direction DirectionValue,
	cell *Cell,
) bool {
	if cell.IsWallSolid(direction) {
		adjacentCell := grid.AdjacentCell(direction, coordinates)
		return adjacentCell != nil && !adjacentCell.Visited
	}
	return false
}

func (grid *Grid) AvailableCellWalls(
	cell *Cell,
	cellCoordinates *Coordinates,
) []DirectionValue {
	response := []DirectionValue{}

	if grid.IsWallAvailable(cellCoordinates, North, cell) {
		response = append(response, North)
	}
	if grid.IsWallAvailable(cellCoordinates, East, cell) {
		response = append(response, East)
	}
	if grid.IsWallAvailable(cellCoordinates, South, cell) {
		response = append(response, South)
	}
	if grid.IsWallAvailable(cellCoordinates, West, cell) {
		response = append(response, West)
	}

	return response
}

func (grid *Grid) UpdateCell(coordinates *Coordinates, cell *Cell) error {
	data, err := grid.CellDataManager.Data(cell)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error updating cell")
	}
	grid.Cells[coordinates.Y][coordinates.X] = data
	return nil
}

func (grid *Grid) CarveCellWall(
	coordinates *Coordinates,
	direction DirectionValue,
) error {
	cell := grid.Cell(coordinates)
	if cell == nil {
		return fmt.Errorf("cell not found")
	}
	cell.CarveWall(direction)
	err := grid.UpdateCell(coordinates, cell)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error updating cell")
	}
	return nil
}
