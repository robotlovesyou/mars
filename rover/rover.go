// package rover contains an implementation of the mars.Rover interface
package rover

import (
	"fmt"

	"github.com/robotlovesyou/mars"
	"github.com/robotlovesyou/mars/position"
)

// Rover contains the rover state and publishes methods for commanding the rover
type Rover struct {
	position   *position.Position
	surfaceMap mars.Map
}

// turn describes a change in direction
type turn struct {
	left  position.Direction
	right position.Direction
}

// deltas describe the change in position given a particular direction
var deltas = map[position.Direction]position.Coordinate{
	position.North: position.NewCoordinate(0, 1),
	position.South: position.NewCoordinate(0, -1),
	position.East:  position.NewCoordinate(1, 0),
	position.West:  position.NewCoordinate(-1, 0),
}

// turns describe the change in direction given a turn instruction and an original direction
var turns = map[position.Direction]turn{
	position.North: {left: position.West, right: position.East},
	position.South: {left: position.East, right: position.West},
	position.East:  {left: position.North, right: position.South},
	position.West:  {left: position.South, right: position.North},
}

// NewPosition creates a new rover initialized at x, y and facing in direction
func New(position *position.Position, surfaceMap mars.Map) *Rover {
	return &Rover{
		position:   position,
		surfaceMap: surfaceMap,
	}
}

// Execute executes a list of instructions and returns the resulting position.
// If an obstacle is encountered it returns the last position and ErrStoppedByObstacle
func (r *Rover) Execute(instructions []mars.Instruction) (*position.Position, error) {
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
	// copy the position to prevent modification of internal state
	return position.NewPosition(r.position.Coordinate().X, r.position.Coordinate().Y, r.position.Direction()), err
}

func (r *Rover) move(instruction mars.Instruction) error {
	scale := 1
	if instruction == mars.Backward {
		scale = -1
	}

	change := deltas[r.position.Direction()].Scale(scale)
	proposed := r.position.Coordinate().Add(change)
	if r.surfaceMap.HasObstacle(proposed) {
		return mars.ErrStoppedByObstacle
	}

	r.position.MoveTo(proposed)
	return nil
}

func (r *Rover) turn(instruction mars.Instruction) {
	change := turns[r.position.Direction()]
	if instruction == mars.Left {
		r.position.TurnTo(change.left)
	} else {
		r.position.TurnTo(change.right)
	}
}
