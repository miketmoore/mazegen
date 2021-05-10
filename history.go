package mazegen

type History struct {
	coordinates []*Coordinates
}

func NewHistory() *History {
	return &History{
		coordinates: []*Coordinates{},
	}
}

func (history *History) Length() int {
	return len(history.coordinates)
}

func (history *History) GetLast() *Coordinates {
	if history.Length() == 0 {
		return nil
	}
	return history.coordinates[history.Length()-1]
}

func (history *History) Push(coordinates *Coordinates) {
	history.coordinates = append(history.coordinates, coordinates)
}

func (history *History) DeleteLast() {
	if history.Length() == 0 {
		return
	}
	history.coordinates = history.coordinates[:history.Length()-1]
}
