package main

import (
	"bufio"
	"bytes"
	"testing"

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
	r.Equal("6, 4, NORTH", out)
	r.NoError(err)
}

func TestRunWithBadConfig(t *testing.T) {
	r := require.New(t)
	rdr := lineReaderFromString("bad" + exampleConfig)
	_, err := run(rdr)
	r.ErrorIs(mars.ErrBadPosition, err)
}
