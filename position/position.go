package position

import "github.com/robotlovesyou/mars"

type Position struct {
	coordinate mars.Coordinate
	direction  mars.Direction
}

func New(x, y int, direction mars.Direction) *Position {
	return &Position{
		coordinate: mars.Coordinate{x, y},
		direction:  direction,
	}
}

func (p *Position) X() int {
	return p.coordinate.X
}

func (p *Position) Y() int {
	return p.coordinate.Y
}

func (p *Position) Direction() mars.Direction {
	return p.direction
}

func (p *Position) Moved(x, y int) mars.Position {
	return &Position{
		coordinate: mars.Coordinate{
			X: p.coordinate.X + x,
			Y: p.coordinate.Y + y,
		},
		direction: p.direction,
	}
}

func (p *Position) Turned(direction mars.Direction) mars.Position {
	return &Position{
		coordinate: p.coordinate,
		direction:  direction,
	}
}
