package rover_test

import (
	"testing"

	"github.com/robotlovesyou/mars/mapping"

	"github.com/robotlovesyou/mars/position"

	"github.com/robotlovesyou/mars"
	"github.com/robotlovesyou/mars/parser"
	"github.com/robotlovesyou/mars/rover"
	"github.com/stretchr/testify/require"
)

// defaultStart returns the default start position.
func defaultStart() position.Position {
	return position.NewPosition(0, 0, position.North)
}

func defaultMap() mars.Map {
	return mapping.New([]position.Coordinate{{1, 4}, {5, 5}, {7, 4}})
}

func instructions(command string, r *require.Assertions) []mars.Instruction {
	instructions, err := parser.Parse(command)
	r.NoError(err)
	return instructions
}

func testRoverExecute(command string, start, expected position.Position, t *testing.T) {
	r := require.New(t)
	rov := rover.New(start, mapping.New(nil))
	inst := instructions(command, r)
	pos, err := rov.Execute(inst)
	r.NoError(err)
	r.Equal(expected, pos)
}

func testRoverExecuteStop(command string, start, expected position.Position, surface mars.Map, t *testing.T) {
	r := require.New(t)
	rov := rover.New(start, surface)
	inst := instructions(command, r)
	pos, err := rov.Execute(inst)
	r.ErrorIs(err, mars.ErrStoppedByObstacle)
	r.Equal(expected, pos)
}

func TestRoverIsCorrectlyInitialized(t *testing.T) {
	testRoverExecute("", defaultStart(), defaultStart(), t)
}

func TestRoverMovesForward(t *testing.T) {
	expected := position.NewPosition(0, 1, position.North)
	testRoverExecute("F", defaultStart(), expected, t)
}

func TestRoverMovesBackward(t *testing.T) {
	expected := position.NewPosition(0, -1, position.North)
	testRoverExecute("B", defaultStart(), expected, t)
}

func TestRoverTurnsLeft(t *testing.T) {
	expected := position.NewPosition(0, 0, position.West)
	testRoverExecute("L", defaultStart(), expected, t)
}

func TestRoverTurnsRight(t *testing.T) {
	expected := position.NewPosition(0, 0, position.East)
	testRoverExecute("R", defaultStart(), expected, t)
}

func TestRoverFollowsInstructions(t *testing.T) {
	start := position.NewPosition(4, 2, position.East)
	expected := position.NewPosition(6, 4, position.North)
	testRoverExecute("FLFFFRFLB", start, expected, t)
}

func TestRoverStopsAtObstacle(t *testing.T) {
	start := position.NewPosition(4, 2, position.East)
	expected := position.NewPosition(5, 4, position.North)
	testRoverExecuteStop("FLFFFRFLB", start, expected, defaultMap(), t)
}

func TestRoverPanicsOnABadInstruction(t *testing.T) {
	r := require.New(t)
	r.Panics(func() {
		rov := rover.New(defaultStart(), defaultMap())
		rov.Execute([]mars.Instruction{"X"})
	})
}
