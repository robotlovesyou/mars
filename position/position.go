// package position contains types which can be used to describe a position (coordinates and direction)
package position

import "fmt"

// Direction describes the direction the rover is facing
type Direction byte

const (
	North Direction = 'N'
	South Direction = 'S'
	East  Direction = 'E'
	West  Direction = 'W'
)

func (d Direction) String() string {
	switch d {
	case North:
		return "NORTH"
	case South:
		return "SOUTH"
	case East:
		return "EAST"
	case West:
		return "WEST"
	default:
		// panics here are caught by the standard library so return a warning
		return fmt.Sprintf("INVALID DIRECTION: %s", string(d))
	}
}

// Coordinate is a point on the surface of mars
type Coordinate struct {
	X int
	Y int
}

func NewCoordinate(x, y int) Coordinate {
	return Coordinate{
		X: x,
		Y: y,
	}
}

// Add two coordinates
func (c Coordinate) Add(addend Coordinate) Coordinate {
	return NewCoordinate(c.X+addend.X, c.Y+addend.Y)
}

// Scale this coordinate by an integer, returning the result
func (c Coordinate) Scale(scalar int) Coordinate {
	return NewCoordinate(c.X*scalar, c.Y*scalar)
}

// Position a position on the surface. It implements the mars.Position interface
type Position struct {
	coordinate Coordinate
	direction  Direction
}

// NewPosition creates a new Position
func NewPosition(x, y int, direction Direction) *Position {
	return &Position{
		coordinate: Coordinate{x, y},
		direction:  direction,
	}
}

func (p *Position) Coordinate() Coordinate {
	return p.coordinate
}

// Direction is the direction of the Position
func (p *Position) Direction() Direction {
	return p.direction
}

// Move returns a new Position which is this position moved by the given amount
func (p *Position) MoveTo(coordinate Coordinate) {
	p.coordinate = coordinate
}

// Turned returns a new Position which is this position turned to the given direction
func (p *Position) TurnTo(direction Direction) {
	p.direction = direction
}
