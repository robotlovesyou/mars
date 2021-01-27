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

type Coordinate struct {
	X int
	Y int
}

// Position describes the x and y coordinates and direction of the rover
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
