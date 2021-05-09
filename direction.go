package mazegen

type DirectionValue string

const (
	North DirectionValue = "north"
	East  DirectionValue = "east"
	South DirectionValue = "south"
	West  DirectionValue = "west"
)

type Direction struct {
	Value DirectionValue
}

func (direction *Direction) Opposite() DirectionValue {
	switch direction.Value {
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
