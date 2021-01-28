package mars

import (
	"errors"

	"github.com/robotlovesyou/mars/position"
)

// Instruction represents a single instruction to the rover
type Instruction string

const (
	Forward  Instruction = "F"
	Backward Instruction = "B"
	Left     Instruction = "L"
	Right    Instruction = "R"
)

var (
	// ErrBadCommands is returned if a command string contains invalid instructions
	ErrBadCommands = errors.New("invalid commands")

	// ErrStoppedByObstacle is returned if the rover stops before hitting an obstacle
	ErrStoppedByObstacle = errors.New("stopped by obstacle")
)

//// Coordinate is a point on the surface of mars
//type Coordinate struct {
//	X int
//	Y int
//}

//// Position describes a point and direction on the surface of mars
//type Position interface {
//	X() int
//	Y() int
//	Direction() Direction
//	Moved(x, y int) Position
//	Turned(to Direction) Position
//}

// Rover interface is implemented by any rover implementation
type Rover interface {
	// Execute carries out a list of instructions and returns the resulting position
	Execute(instructions []Instruction) (position.Position, error)
}

// Map describes a map of the surface and can be queried for obstacles
type Map interface {
	HasObstacle(x, y int) bool
}
