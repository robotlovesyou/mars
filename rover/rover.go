// package rover contains an implementation of the mars.Rover interface
package rover

import (
	"fmt"

	"github.com/robotlovesyou/mars"
	"github.com/robotlovesyou/mars/position"
)

// Rover contains the rover state and publishes methods for commanding the rover
type Rover struct {
	position   position.Position
	surfaceMap mars.Map
}

// delta describes a change in position
type delta struct {
	x int
	y int
}

// turn describes a change in direction
type turn struct {
	left  position.Direction
	right position.Direction
}

// deltas describe the change in position given a particular direction
var deltas = map[position.Direction]delta{
	position.North: {x: 0, y: 1},
	position.South: {x: 0, y: -1},
	position.East:  {x: 1, y: 0},
	position.West:  {x: -1, y: 0},
}

// turns describe the change in direction given a turn instruction and an original direction
var turns = map[position.Direction]turn{
	position.North: {left: position.West, right: position.East},
	position.South: {left: position.East, right: position.West},
	position.East:  {left: position.North, right: position.South},
	position.West:  {left: position.South, right: position.North},
}

// NewPosition creates a new rover initialized at x, y and facing in direction
func New(position position.Position, surfaceMap mars.Map) *Rover {
	return &Rover{
		position:   position,
		surfaceMap: surfaceMap,
	}
}

// Execute executes a list of instructions and returns the resulting position.
// If an obstacle is encountered it returns the last position and ErrStoppedByObstacle
func (r *Rover) Execute(instructions []mars.Instruction) (position.Position, error) {
	var err error
	for _, instruction := range instructions {
		switch instruction {
		case mars.Forward, mars.Backward:
			err = r.move(instruction)
		case mars.Left, mars.Right:
			r.turn(instruction)
		default:
			panic(fmt.Sprintf("bad instruction %s", instruction))
		}
		if err != nil {
			break
		}

	}
	return r.position, err
}

func (r *Rover) move(instruction mars.Instruction) error {
	scale := 1
	if instruction == mars.Backward {
		scale = -1
	}

	change := deltas[r.position.Direction()]
	proposed := r.position.Moved(change.x*scale, change.y*scale)
	if r.surfaceMap.HasObstacle(position.Coordinate{proposed.X(), proposed.Y()}) {
		return mars.ErrStoppedByObstacle
	}

	r.position = proposed
	return nil
}

func (r *Rover) turn(instruction mars.Instruction) {
	change := turns[r.position.Direction()]
	if instruction == mars.Left {
		r.position = r.position.Turned(change.left)
	} else {
		r.position = r.position.Turned(change.right)
	}
}
