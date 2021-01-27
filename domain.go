package mars

import "errors"

// Direction describes the direction the rover is facing
type Direction string

// Instruction represents a single instruction to the rover
type Instruction string

const (
	North    Direction   = "NORTH"
	South    Direction   = "SOUTH"
	East     Direction   = "EAST"
	West     Direction   = "WEST"
	Forward  Instruction = "F"
	Backward Instruction = "B"
	Left     Instruction = "L"
	Right    Instruction = "R"
)

// ErrBadCommands is returned if a command string contains invalid instructions
var ErrBadCommands = errors.New("invalid commands")

// Coordinate is a point on the surface of mars
type Coordinate struct {
	X int
	Y int
}

// Position describes a point and direction on the surface of mars
type Position interface {
	X() int
	Y() int
	Direction() Direction
	Moved(x, y int) Position
	Turned(to Direction) Position
}

// Rover interface is implemented by any rover implementation
type Rover interface {
	// Execute carries out a list of instructions and returns the resulting position
	Execute(instructions []Instruction) Position
}

// Map describes a map of the surface and can be queried for obstacles
type Map interface {
	HasObstacle(x, y int) bool
}
