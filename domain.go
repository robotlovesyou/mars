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
	// ErrBadPosition is returned if a parsed position is invalid
	ErrBadPosition = errors.New("bad position")
	// ErrBadCoordinate is returned if a parsed coordinate is invalid
	ErrBadCoordinate = errors.New("bad coordinate")
)

// Rover interface is implemented by any rover implementation
type Rover interface {
	// Execute carries out a list of instructions and returns the resulting position
	Execute(instructions []Instruction) (position.Position, error)
}

// Map describes a map of the surface and can be queried for obstacles
type Map interface {
	HasObstacle(coordinate position.Coordinate) bool
}
