package mazegen

type Coordinates struct {
	X, Y int
}

func NewCoordinates(x, y int) *Coordinates {
	return &Coordinates{X: x, Y: y}
}
