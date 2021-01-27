package rover_test

import (
	"testing"

	"github.com/robotlovesyou/mars/position"

	"github.com/robotlovesyou/mars"
	"github.com/robotlovesyou/mars/parser"
	"github.com/robotlovesyou/mars/rover"
	"github.com/stretchr/testify/require"
)

// defaultStart returns the default start position.
func defaultStart() mars.Position {
	return position.New(0, 0, mars.North)
}

func instructions(command string, r *require.Assertions) []mars.Instruction {
	instructions, err := parser.Parse(command)
	r.NoError(err)
	return instructions
}

func testRoverExecute(command string, start, expected mars.Position, t *testing.T) {
	r := require.New(t)
	rov := rover.New(start)
	inst := instructions(command, r)
	pos := rov.Execute(inst)
	r.Equal(expected, pos)
}

func TestRoverIsCorrectlyInitialized(t *testing.T) {
	testRoverExecute("", defaultStart(), defaultStart(), t)
}

func TestRoverMovesForward(t *testing.T) {
	expected := position.New(0, 1, mars.North)
	testRoverExecute("F", defaultStart(), expected, t)
}

func TestRoverMovesBackward(t *testing.T) {
	expected := position.New(0, -1, mars.North)
	testRoverExecute("B", defaultStart(), expected, t)
}

func TestRoverTurnsLeft(t *testing.T) {
	expected := position.New(0, 0, mars.West)
	testRoverExecute("L", defaultStart(), expected, t)
}

func TestRoverTurnsRight(t *testing.T) {
	expected := position.New(0, 0, mars.East)
	testRoverExecute("R", defaultStart(), expected, t)
}

func TestRoverFollowsInstructions(t *testing.T) {
	start := position.New(4, 2, mars.East)
	expected := position.New(6, 4, mars.North)
	testRoverExecute("FLFFFRFLB", start, expected, t)
}
