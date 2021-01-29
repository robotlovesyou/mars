package parser_test

import (
	"testing"

	"github.com/robotlovesyou/mars/position"

	"github.com/robotlovesyou/mars"
	"github.com/robotlovesyou/mars/parser"
	"github.com/stretchr/testify/require"
)

func TestCorrectlyParsesCommand(t *testing.T) {
	r := require.New(t)
	instructions, err := parser.ParseCommands("FLFFFRFLB")
	r.NoError(err)
	r.Equal([]mars.Instruction{
		mars.Forward,
		mars.Left,
		mars.Forward,
		mars.Forward,
		mars.Forward,
		mars.Right,
		mars.Forward,
		mars.Left,
		mars.Backward}, instructions)

}

func TestFailsOnBadCommand(t *testing.T) {
	r := require.New(t)
	instructions, err := parser.ParseCommands("FLFFRFX")
	r.ErrorIs(err, mars.ErrBadCommands)
	r.Nil(instructions)
}

func TestReturnsEmptyInstructionsOnEmptyCommand(t *testing.T) {
	r := require.New(t)
	instructions, err := parser.ParseCommands("")
	r.NoError(err)
	r.Len(instructions, 0)
}

func TestCorrectlyParsesPosition(t *testing.T) {
	tests := []struct {
		x int
		y int
		d position.Direction
		s string
	}{
		{4, -2, position.East, "4, -2, EAST"},
		{-44, 22, position.West, "-44, 22, WEST"},
		{444, -222, position.North, "444, -222, NORTH"},
		{-4444, 2222, position.South, "-4444, 2222, SOUTH"},
	}
	r := require.New(t)
	for _, test := range tests {
		pos, err := parser.ParsePosition(test.s)
		r.NoError(err)
		r.Equal(position.NewPosition(position.NewCoordinate(test.x, test.y), test.d), pos, test.s)
	}

}

func TestReturnsBadPositionForBadPositionString(t *testing.T) {
	tests := []string{"a, 2, EAST", "2, a, EAST", "1, 2, 3, WEST", "1, WEST", "WEST", "4, 2, BRIAN"}
	r := require.New(t)
	for _, test := range tests {
		_, err := parser.ParsePosition(test)
		r.ErrorIs(err, mars.ErrBadPosition, test)
	}
}

func TestParsesCoordinates(t *testing.T) {
	tests := []struct {
		x int
		y int
		s string
	}{
		{4, -2, "4, -2"},
		{-44, 22, "-44, 22"},
	}

	r := require.New(t)
	for _, test := range tests {
		coord, err := parser.ParseCoordinate(test.s)
		r.NoError(err)
		r.Equal(position.NewCoordinate(test.x, test.y), coord, test.s)
	}
}

func TestReturnsBadCoordinateForBadCoordinateString(t *testing.T) {
	tests := []string{"1, 2, 3", "1", "a, 2", "2, a"}
	r := require.New(t)
	for _, test := range tests {
		_, err := parser.ParseCoordinate(test)
		r.Equal(mars.ErrBadCoordinate, err)
	}
}
