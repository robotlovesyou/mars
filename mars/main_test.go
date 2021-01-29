package main

import (
	"bufio"
	"bytes"
	"errors"
	"testing"

	"github.com/robotlovesyou/mars/position"

	"github.com/robotlovesyou/mars"

	"github.com/stretchr/testify/require"
)

func lineReaderFromString(s string) *bufio.Reader {
	return bufio.NewReader(bytes.NewBufferString(s))
}

func TestRun(t *testing.T) {
	r := require.New(t)
	rdr := lineReaderFromString(exampleConfig)
	out, err := run(rdr)
	r.NoError(err)
	r.Equal(position.NewPosition(position.NewCoordinate(6, 4), position.North), out)
}

func TestPrepareOutputOnSuccess(t *testing.T) {
	r := require.New(t)
	pos := position.NewPosition(position.NewCoordinate(6, 4), position.North)
	r.Equal("6, 4, NORTH", prepareOutput(pos, nil))
}

func TestPrepareOutputStopped(t *testing.T) {
	r := require.New(t)
	pos := position.NewPosition(position.NewCoordinate(6, 4), position.North)
	r.Equal("6, 4, NORTH STOPPED", prepareOutput(pos, mars.ErrStoppedByObstacle))
}

func TestPrepareOutputUnexpected(t *testing.T) {
	r := require.New(t)
	r.Equal("encountered unexpected error: nope", prepareOutput(nil, errors.New("nope")))
}

func TestRunWithBadConfig(t *testing.T) {
	r := require.New(t)
	rdr := lineReaderFromString("bad" + exampleConfig)
	_, err := run(rdr)
	r.ErrorIs(mars.ErrBadPosition, err)
}
