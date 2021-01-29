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

func testPosition() *position.Position {
	return position.NewPosition(position.NewCoordinate(1, 2), position.North)
}

func TestCanAddACoordinateToACoordinate(t *testing.T) {
	r := require.New(t)
	res := position.NewCoordinate(1, 2).Add(position.NewCoordinate(1, 2))
	r.Equal(position.NewCoordinate(2, 4), res)
}

func TestCanScaleACoordinate(t *testing.T) {
	r := require.New(t)
	res := position.NewCoordinate(1, 2).Scale(-1)
	r.Equal(position.NewCoordinate(-1, -2), res)
}

func TestPositionReportsCorrectCoordinate(t *testing.T) {
	r := require.New(t)
	pos := testPosition()
	r.Equal(position.NewCoordinate(1, 2), pos.Coordinate())
}

func TestPositionReportsCorrectDirection(t *testing.T) {
	r := require.New(t)
	pos := testPosition()
	r.Equal(position.North, pos.Direction())
}

func TestMovesToExpectedCoordinate(t *testing.T) {
	r := require.New(t)
	pos := testPosition()
	to := position.NewCoordinate(1, 1)
	pos.MoveTo(to)
	r.Equal(to, pos.Coordinate())
}

func TestTurnsToExpectedDirection(t *testing.T) {
	r := require.New(t)
	pos := testPosition()
	pos.TurnTo(position.West)
	r.Equal(position.West, pos.Direction())
}
