package rover_test

import (
	"github.com/robotlovesyou/mars"
	"github.com/robotlovesyou/mars/parser"
	"github.com/robotlovesyou/mars/rover"
	"github.com/stretchr/testify/require"
	"testing"
)

func newRover() *rover.Rover {
	return rover.New(0, 0, mars.North)
}

func instructions(command string, r *require.Assertions) []mars.Instruction {
	instructions, err := parser.Parse(command)
	r.NoError(err)
	return instructions
}

func TestRoverIsCorrectlyInitialized(t *testing.T) {
	rov := newRover()
	r := require.New(t)
	inst := instructions("", r)
	pos := rov.Execute(inst)
	r.Equal(mars.Position{
		X:         0,
		Y:         0,
		Direction: mars.North,
	}, pos)
}

func TestRoverMovesForward(t *testing.T) {
	rov := newRover()
	r := require.New(t)
	inst := instructions("F", r)
	pos := rov.Execute(inst)
	r.Equal(mars.Position{
		X:         0,
		Y:         1,
		Direction: mars.North,
	}, pos)
}

func TestRoverMovesBackward(t *testing.T) {
	rov := newRover()
	r := require.New(t)
	inst := instructions("B", r)
	pos := rov.Execute(inst)
	r.Equal(mars.Position{
		X:         0,
		Y:         -1,
		Direction: mars.North,
	}, pos)
}

func TestRoverTurnsLeft(t *testing.T) {
	rov := newRover()
	r := require.New(t)
	inst := instructions("L", r)
	pos := rov.Execute(inst)
	r.Equal(mars.Position{
		X:         0,
		Y:         0,
		Direction: mars.West,
	}, pos)
}

func TestRoverTurnsRight(t *testing.T) {
	rov := newRover()
	r := require.New(t)
	inst := instructions("R", r)
	pos := rov.Execute(inst)
	r.Equal(mars.Position{
		X:         0,
		Y:         0,
		Direction: mars.East,
	}, pos)
}

func TestRoverFollowsInstructions(t *testing.T) {
	rov := rover.New(4, 2, mars.East)
	r := require.New(t)
	inst := instructions("FLFFFRFLB", r)
	pos := rov.Execute(inst)
	r.Equal(mars.Position{
		X:         6,
		Y:         4,
		Direction: mars.North,
	}, pos)
}
