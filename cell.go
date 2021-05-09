package mazegen

import (
	"fmt"	"strings"
)


type Cell struct {
	Visited bool
	Walls   map[DirectionValue]bool
}

func NewCell() *Cell {
	return &Cell{
		Visited: false,
		Walls: map[DirectionValue]bool{
			North: true,
			East:  true,
			South: true,
			West:  true,
		},
	}
}

func NewFromData(data string) *Cell {
	// const [visited, north, east, south, west] = data.split('')
	// const cell = Cell.new()
	// if (north === '0') cell.carveWall('north')
	// if (east === '0') cell.carveWall('east')
	// if (west === '0') cell.carveWall('west')
	// if (south === '0') cell.carveWall('south')
	// return cell
	pieces := strings.Split(data, "")
	visited := pieces[0]
	north := pieces[1]
	east := pieces[2]
	south := pieces[3]
	west := pieces[4]

	cell := NewCell()
	if north == "0" {
		cell.CarveWall(North  )
	}
}

func (cell *Cell) CarveWall(direction DirectionValue) {
	cell.Walls[direction] = false
	cell.Visited = true
}

func (cell *Cell) IsWallSolid(direction DirectionValue) bool {
	return cell.Walls[direction] == true
}

func (cell *Cell) OppositeDirection(direction DirectionValue) DirectionValue {
	switch direction {
	case North:
		return South
	case East:
		return West
	case South:
		return North
	case West:
		return East
	}

	return "invalidDirection"
}

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
