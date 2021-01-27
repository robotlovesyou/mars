package rover

import (
	"fmt"
	"github.com/robotlovesyou/mars"
)

type Rover struct {
	position mars.Position
}

type delta struct {
	x int
	y int
}

type facing struct {
	left mars.Direction
	right mars.Direction
}

var deltas = map[mars.Direction]delta{
	mars.North: {x:0,  y:1},
	mars.South: {x: 0, y: -1},
	mars.East: {x: 1, y: 0},
	mars.West: {x: -1, y: 0},
}

var facings = map[mars.Direction]facing {
	mars.North: {left: mars.West, right: mars.East},
	mars.South: {left: mars.East, right: mars.West},
	mars.East: {left: mars.North, right: mars.South},
	mars.West: {left: mars.South, right: mars.North},
}

func New(x, y int, direction mars.Direction) *Rover {
	return &Rover {
		position: mars.Position{
			X: x,
			Y: y,
			Direction: direction,
		},
	}
}

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
	change := deltas[r.position.Direction]
	r.position.X += change.x * scale
	r.position.Y += change.y * scale
}

func (r *Rover) turn(instruction mars.Instruction) {
	change := facings[r.position.Direction]
	if instruction == mars.Left {
		r.position.Direction = change.left
	} else {
		r.position.Direction = change.right
	}
}