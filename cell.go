package mazegen

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
