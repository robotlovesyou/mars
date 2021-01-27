// package position contains types which can be used to describe a position (coordinates and direction)
package position

import "github.com/robotlovesyou/mars"

// Position a position on the surface. It implements the mars.Position interface
type Position struct {
	coordinate mars.Coordinate
	direction  mars.Direction
}

// New creates a new Position
func New(x, y int, direction mars.Direction) *Position {
	return &Position{
		coordinate: mars.Coordinate{x, y},
		direction:  direction,
	}
}

// X is the x coordinate of the position
func (p *Position) X() int {
	return p.coordinate.X
}

// Y is the y coordinate of the position
func (p *Position) Y() int {
	return p.coordinate.Y
}

// Direction is the direction of the Position
func (p *Position) Direction() mars.Direction {
	return p.direction
}

// Move returns a new Position which is this position moved by the given amount
func (p *Position) Moved(x, y int) mars.Position {
	return &Position{
		coordinate: mars.Coordinate{
			X: p.coordinate.X + x,
			Y: p.coordinate.Y + y,
		},
		direction: p.direction,
	}
}

// Turned returns a new Position which is this position turned to the given direction
func (p *Position) Turned(direction mars.Direction) mars.Position {
	return &Position{
		coordinate: p.coordinate,
		direction:  direction,
	}
}
