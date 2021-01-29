// package position contains types which can be used to describe a position on the surface (coordinates and direction)
package position

import "fmt"

// Direction describes the direction the rover is facing
type Direction byte

// Direction Constants
const (
	North Direction = 'N'
	South Direction = 'S'
	East  Direction = 'E'
	West  Direction = 'W'
)

// String interface implementation for Direction
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

// NewCoordinate constructs a new coordinate
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

// String interface implamentation for coordinate
func (c Coordinate) String() string {
	return fmt.Sprintf("%d, %d", c.X, c.Y)
}

// Position a position on the surface. It implements the mars.Position interface
type Position struct {
	coordinate Coordinate
	direction  Direction
}

// NewPosition creates a new Position
func NewPosition(coordinate Coordinate, direction Direction) *Position {
	return &Position{
		coordinate: coordinate,
		direction:  direction,
	}
}

// Coordinate of this position
func (p *Position) Coordinate() Coordinate {
	return p.coordinate
}

// Direction of the Position
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

// String interface implementation for Position
func (p *Position) String() string {
	return fmt.Sprintf("%v, %v", p.coordinate, p.direction)
}
