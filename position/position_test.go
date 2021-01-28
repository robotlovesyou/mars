package position_test

import (
	"fmt"
	"testing"

	"github.com/robotlovesyou/mars/position"

	"github.com/stretchr/testify/require"
)

func TestDirectionStringIsCorrect(t *testing.T) {
	r := require.New(t)
	r.Equal("NORTH", fmt.Sprintf("%v", position.North))
	r.Equal("SOUTH", fmt.Sprintf("%v", position.South))
	r.Equal("EAST", fmt.Sprintf("%v", position.East))
	r.Equal("WEST", fmt.Sprintf("%v", position.West))
	r.Equal("INVALID DIRECTION: X", fmt.Sprintf("%v", position.Direction('X')))

}

func testPosition() position.Position {
	return position.NewPosition(1, 2, position.North)
}

func TestCanAddACoordinateToACoordinate(t *testing.T) {
	r := require.New(t)
	res := position.Coordinate{1, 2}.Add(position.Coordinate{1, 2})
	r.Equal(position.Coordinate{2, 4}, res)
}

func TestCanScaleACoordinate(t *testing.T) {
	r := require.New(t)
	res := position.Coordinate{1, 2}.Scale(-1)
	r.Equal(position.Coordinate{-1, -2}, res)
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
