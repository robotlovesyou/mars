// package rover contains an implementation of the mars.Rover interface
package rover

import (
	"fmt"

	"github.com/robotlovesyou/mars"
)

// Rover contains the rover state and publishes methods for commanding the rover
type Rover struct {
	position mars.Position
}

// delta describes a change in position
type delta struct {
	x int
	y int
}

// turn describes a change in direction
type turn struct {
	left  mars.Direction
	right mars.Direction
}

// deltas describe the change in position given a particular direction
var deltas = map[mars.Direction]delta{
	mars.North: {x: 0, y: 1},
	mars.South: {x: 0, y: -1},
	mars.East:  {x: 1, y: 0},
	mars.West:  {x: -1, y: 0},
}

// turns describe the change in direction given a turn instruction and an original direction
var turns = map[mars.Direction]turn{
	mars.North: {left: mars.West, right: mars.East},
	mars.South: {left: mars.East, right: mars.West},
	mars.East:  {left: mars.North, right: mars.South},
	mars.West:  {left: mars.South, right: mars.North},
}

// New creates a new rover initialized at x, y and facing in direction
func New(position mars.Position) *Rover {
	return &Rover{
		position,
	}
}

// Execute executes a list of instructions and returns the resulting position
func (r *Rover) Execute(instructions []mars.Instruction) mars.Position {
	for _, instruction := range instructions {
		switch instruction {
		case mars.Forward, mars.Backward:
			r.move(instruction)
		case mars.Left, mars.Right:
			r.turn(instruction)
		default:
			panic(fmt.Sprintf("bad instruction %s", instruction))
		}

	}
	return r.position
}

func (r *Rover) move(instruction mars.Instruction) {
	scale := 1
	if instruction == mars.Backward {
		scale = -1
	}
	change := deltas[r.position.Direction()]
	r.position = r.position.Moved(change.x*scale, change.y*scale)
}

func (r *Rover) turn(instruction mars.Instruction) {
	change := turns[r.position.Direction()]
	if instruction == mars.Left {
		r.position = r.position.Turned(change.left)
	} else {
		r.position = r.position.Turned(change.right)
	}
}
