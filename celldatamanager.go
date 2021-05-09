package mazegen

import (
	"fmt"
	"strings"
)

// north east south west
// 1 = solid
// 0 = carved
var wallCombinations = [16]string{
	"0000",
	"0001",
	"0010",
	"0100",
	"1000",
	"1111",
	"0011",
	"0110",
	"1100",
	"1010",
	"0101",
	"1110",
	"0111",
	"1011",
	"1101",
	"1001",
}

func buildCellCombinations() [32]string {
	cellCombinations := [32]string{}

	// Visited
	for comboIndex, combo := range wallCombinations {
		cellCombinations[comboIndex] = fmt.Sprintf("1%s", combo)
	}
	// Not visited
	for comboIndex, combo := range wallCombinations {
		cellCombinations[comboIndex] = fmt.Sprintf("0%s", combo)
	}

	return cellCombinations
}

type CellDataManager struct {
	CellCombinations [32]string
}

func NewCellDataManager() *CellDataManager {
	return &CellDataManager{
		CellCombinations: buildCellCombinations(),
	}
}

func (mgr *CellDataManager) NewFromData(data string) *Cell {
	// const [visited, north, east, south, west] = data.split('')
	// const cell = Cell.new()
	// if (north === '0') cell.carveWall('north')
	// if (east === '0') cell.carveWall('east')
	// if (west === '0') cell.carveWall('west')
	// if (south === '0') cell.carveWall('south')
	// return cell
	pieces := strings.Split(data, "")
	// visited := pieces[0]
	north := pieces[1]
	east := pieces[2]
	south := pieces[3]
	west := pieces[4]

	cell := NewCell()
	if north == "0" {
		cell.CarveWall(North)
	}
	if east == "0" {
		cell.CarveWall(East)
	}
	if west == "0" {
		cell.CarveWall(West)
	}
	if south == "0" {
		cell.CarveWall(South)
	}
	return cell
}

func (mgr *CellDataManager) Data(cell *Cell) (string, error) {
	var visited string
	if cell.Visited {
		visited = "1"
	} else {
		visited = "2"
	}

	var north string
	if cell.Walls[North] {
		north = "1"
	} else {
		north = "2"
	}

	var east string
	if cell.Walls[East] {
		east = "1"
	} else {
		east = "2"
	}

	var south string
	if cell.Walls[South] {
		south = "1"
	} else {
		south = "2"
	}

	var west string
	if cell.Walls[West] {
		west = "1"
	} else {
		west = "2"
	}

	current := fmt.Sprintf("%s%s%s%s%s", visited, north, east, south, west)

	matchFound := false
	for _, combo := range mgr.CellCombinations {
		if current == combo {
			matchFound = true
			break
		}
	}
	if !matchFound {
		return "", fmt.Errorf("match not found for value=%s", current)
	}

	return current, nil
}
