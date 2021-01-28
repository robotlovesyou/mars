package position_test

import (
	"testing"

	"github.com/robotlovesyou/mars/position"

	"github.com/stretchr/testify/require"
)

func testPosition() position.Position {
	return position.New(1, 2, position.North)
}

func TestPositionReportsCorrectX(t *testing.T) {
	r := require.New(t)
	pos := testPosition()
	r.Equal(1, pos.X())
}

func TestPositionReportsCorrectY(t *testing.T) {
	r := require.New(t)
	pos := testPosition()
	r.Equal(2, pos.Y())
}

func TestPositionReportsCorrectDirection(t *testing.T) {
	r := require.New(t)
	pos := testPosition()
	r.Equal(position.North, pos.Direction())
}

func TestMovesToExpectedCoordinate(t *testing.T) {
	r := require.New(t)
	pos := testPosition()
	moved := pos.Moved(1, 1)
	r.Equal(2, moved.X())
	r.Equal(3, moved.Y())
	r.Equal(pos.Direction(), moved.Direction())
}

func TestTurnsToExpectedDirection(t *testing.T) {
	r := require.New(t)
	pos := testPosition()
	turned := pos.Turned(position.West)
	r.Equal(pos.X(), turned.X())
	r.Equal(pos.Y(), turned.Y())
	r.Equal(position.West, turned.Direction())
}
