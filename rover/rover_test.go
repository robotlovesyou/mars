package rover_test

import (
	"testing"

	"github.com/robotlovesyou/mars"
	"github.com/robotlovesyou/mars/parser"
	"github.com/robotlovesyou/mars/rover"
	"github.com/stretchr/testify/require"
)

// defaultStart returns the default start position.
func defaultStart() mars.Position {
	return mars.Position{
		X:         0,
		Y:         0,
		Direction: mars.North,
	}
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
	expected := mars.Position{
		X:         0,
		Y:         1,
		Direction: mars.North,
	}
	testRoverExecute("F", defaultStart(), expected, t)
}

func TestRoverMovesBackward(t *testing.T) {
	expected := mars.Position{
		X:         0,
		Y:         -1,
		Direction: mars.North,
	}
	testRoverExecute("B", defaultStart(), expected, t)
}

func TestRoverTurnsLeft(t *testing.T) {
	expected := mars.Position{
		X:         0,
		Y:         0,
		Direction: mars.West,
	}
	testRoverExecute("L", defaultStart(), expected, t)
}

func TestRoverTurnsRight(t *testing.T) {
	expected := mars.Position{
		X:         0,
		Y:         0,
		Direction: mars.East,
	}
	testRoverExecute("R", defaultStart(), expected, t)
}

func TestRoverFollowsInstructions(t *testing.T) {
	start := mars.Position{
		X:         4,
		Y:         2,
		Direction: mars.East,
	}
	expected := mars.Position{
		X:         6,
		Y:         4,
		Direction: mars.North,
	}
	testRoverExecute("FLFFFRFLB", start, expected, t)
}
