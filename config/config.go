package config

import (
	"bufio"
	"io"
	"strings"

	"github.com/robotlovesyou/mars/mapping"
	"github.com/robotlovesyou/mars/parser"

	"github.com/robotlovesyou/mars"
	"github.com/robotlovesyou/mars/position"
)

type Config struct {
	Start        *position.Position
	Instructions []mars.Instruction
	Map          mars.Map
}

func Load(reader *bufio.Reader) (*Config, error) {

	startLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, mars.ErrBadPosition
	}

	start, err := parser.ParsePosition(strings.TrimSpace(startLine))
	if err != nil {
		return nil, err
	}

	commandLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, mars.ErrBadCommands
	}

	instructions, err := parser.ParseCommands(strings.TrimSpace(commandLine))
	if err != nil {
		return nil, err
	}

	var coord position.Coordinate
	var coordLine string
	coords := make([]position.Coordinate, 0)

	for coordLine, err = reader.ReadString('\n'); err == nil; coordLine, err = reader.ReadString('\n') {
		coordLine = strings.TrimSpace(coordLine)
		if len(coordLine) == 0 {
			break
		}

		coord, err = parser.ParseCoordinate(coordLine)
		if err != nil {
			return nil, err
		}
		coords = append(coords, coord)
	}
	if err != nil && err != io.EOF {
		return nil, err
	}

	surface := mapping.New(coords)
	return &Config{
		Start:        start,
		Instructions: instructions,
		Map:          surface,
	}, nil
}
