package mazegen

import (
	"fmt"
	"math/rand"
)

func BuildMaze(rows, columns int, random *rand.Rand) (*Grid, error) {
	grid, err := NewGrid(rows, columns, random)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error building new grid")
	}

	// history := []*Coordinates{grid.RandomCoordinates()}
	history := NewHistory()
	history.Push(grid.RandomCoordinates())

	running := true

	for running {
		// coordinates := history[len(history)-1]
		coordinates := history.GetLast()
		cell := grid.Cell(coordinates)
		if cell == nil {
			return nil, fmt.Errorf("cell not found")
		}

		availableWalls := grid.AvailableCellWalls(cell, coordinates)

		if len(availableWalls) == 0 {
			if history.Length() >= 2 {
				// TODO pop history
				history.DeleteLast()
			} else {
				running = false
			}
		} else {
			wallIndex := random.Intn(len(availableWalls))
			availableWall := availableWalls[wallIndex]
			grid.CarveCellWall(coordinates, availableWall)

			adjacentCoordinates := grid.AdjacentCoordinates(availableWall, coordinates)
			if adjacentCoordinates != nil {
				adjacentCell := grid.Cell(adjacentCoordinates)
				if adjacentCell != nil && !adjacentCell.Visited {
					oppositeDirection := adjacentCell.OppositeDirection(availableWall)
					grid.CarveCellWall(adjacentCoordinates, oppositeDirection)
					// history = append(history, adjacentCoordinates)
					history.Push(adjacentCoordinates)
				}
			}
		}

	}

	return grid, nil
}
