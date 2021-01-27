package parser_test

import (
	"github.com/robotlovesyou/mars"
	"github.com/robotlovesyou/mars/parser"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCorrectlyParsesCommand(t *testing.T) {
	r := require.New(t)
	instructions, err := parser.Parse("FLFFFRFLB")
	r.NoError(err)
	r.Equal([]mars.Instruction{"F", "L", "F", "F" , "F", "R", "F", "L", "B"}, instructions)

}

func TestFailsOnBadCommand(t *testing.T) {
	r := require.New(t)
	instructions, err := parser.Parse("FLFFRFX")
	r.ErrorIs(err, mars.ErrBadCommands)
	r.Nil(instructions)
}

func TestReturnsEmptyInstructionsOnEmptyCommand(t *testing.T) {
	r := require.New(t)
	instructions, err := parser.Parse("")
	r.NoError(err)
	r.Len(instructions, 0)
}
