package config_test

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/robotlovesyou/mars"
	"github.com/robotlovesyou/mars/config"
	"github.com/robotlovesyou/mars/position"
	"github.com/stretchr/testify/require"
)

const goodConfig = `-4, 2, EAST
FLFFFRFLB
1, 4
3, 5
-7, 4
`

const badStart = `4, 2, 3, EAST
FLFFFRFLB
1, 4
3, 5
7, 4
`

const badCommand = `4, 2, EAST
XFLFFFRFLB
1, 4
3, 5
7, 4
`

const badMap = `4, 2, EAST
FLFFFRFLB
1, 4
3, 5
a, 7, 4
`

const missingInstructions = "4, 2, EAST\n"

const noMap = `4, 2, EAST
FLFFFRFLB
`

func lineReaderWith(s string) *bufio.Reader {
	buffer := bytes.NewBufferString(s)
	return bufio.NewReader(buffer)
}

func TestReadsConfig(t *testing.T) {
	r := require.New(t)
	rdr := lineReaderWith(goodConfig)
	conf, err := config.Load(rdr)
	r.NoError(err)
	r.Equal(position.NewPosition(position.NewCoordinate(-4, 2), position.East), conf.Start)
	r.Equal([]mars.Instruction{
		mars.Forward,
		mars.Left,
		mars.Forward,
		mars.Forward,
		mars.Forward,
		mars.Right,
		mars.Forward,
		mars.Left,
		mars.Backward}, conf.Instructions)
	r.True(conf.Map.HasObstacle(position.NewCoordinate(1, 4)), "1, 4")
	r.True(conf.Map.HasObstacle(position.NewCoordinate(3, 5)), "3, 5")
	r.True(conf.Map.HasObstacle(position.NewCoordinate(-7, 4)), "-7, 4")
	r.False(conf.Map.HasObstacle(position.NewCoordinate(4, 2)), "4, 2 should have no obstacle")
}

func TestReturnsBadStartForBadStartData(t *testing.T) {
	r := require.New(t)
	rdr := lineReaderWith(badStart)
	_, err := config.Load(rdr)
	r.ErrorIs(mars.ErrBadPosition, err)
}

func TestReturnsBadStartForMissingStartData(t *testing.T) {
	r := require.New(t)
	rdr := bufio.NewReader(bytes.NewBuffer(nil))
	_, err := config.Load(rdr)
	r.ErrorIs(mars.ErrBadPosition, err)
}

func TestReturnsBadCommandForBadCommandData(t *testing.T) {
	r := require.New(t)
	rdr := lineReaderWith(badCommand)
	_, err := config.Load(rdr)
	r.ErrorIs(mars.ErrBadCommands, err)
}

func TestReturnsBadCommandForMissingCommandData(t *testing.T) {
	r := require.New(t)
	rdr := lineReaderWith(missingInstructions)
	_, err := config.Load(rdr)
	r.ErrorIs(mars.ErrBadCommands, err)
}

func TestReturnsBadMapForBadMapData(t *testing.T) {
	r := require.New(t)
	rdr := lineReaderWith(badMap)
	_, err := config.Load(rdr)
	r.ErrorIs(mars.ErrBadCoordinate, err)
}

func TestReturnsOKWithNoMap(t *testing.T) {
	r := require.New(t)
	rdr := lineReaderWith(noMap)
	conf, err := config.Load(rdr)
	r.NotNil(conf)
	r.NoError(err)
}
