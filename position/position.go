// package position contains types which can be used to describe a position (coordinates and direction)
package position

// Direction describes the direction the rover is facing
type Direction string

const (
	North Direction = "NORTH"
	South Direction = "SOUTH"
	East  Direction = "EAST"
	West  Direction = "WEST"
)

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
func NewPosition(x, y int, direction Direction) Position {
	return Position{
		coordinate: Coordinate{x, y},
		direction:  direction,
	}
}

// X is the x coordinate of the position
func (p Position) X() int {
	return p.coordinate.X
}

// Y is the y coordinate of the position
func (p Position) Y() int {
	return p.coordinate.Y
}

// Direction is the direction of the Position
func (p Position) Direction() Direction {
	return p.direction
}

// Move returns a new Position which is this position moved by the given amount
func (p Position) Moved(x, y int) Position {
	return Position{
		coordinate: Coordinate{
			X: p.coordinate.X + x,
			Y: p.coordinate.Y + y,
		},
		direction: p.direction,
	}
}

// Turned returns a new Position which is this position turned to the given direction
func (p Position) Turned(direction Direction) Position {
	return Position{
		coordinate: p.coordinate,
		direction:  direction,
	}
}
