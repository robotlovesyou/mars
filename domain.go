package mars

import "errors"

type Direction string
type Instruction string

const (
	North Direction = "NORTH"
	South Direction = "SOUTH"
	East Direction = "EAST"
	West Direction = "WEST"
)

var ErrBadCommands = errors.New("invalid commands")

type Position struct {
	X int
	Y int
	Direction Direction
}

type Rover interface {
	Execute(instructions []Instruction) Position
}
