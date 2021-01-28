package parser

import (
	"math/bits"
	"regexp"
	"strconv"

	mars "github.com/robotlovesyou/mars"
	"github.com/robotlovesyou/mars/position"
)

var (
	positionRegex   = regexp.MustCompile(`^(\d+), (\d+), (NORTH|SOUTH|EAST|WEST)$`)
	coordinateRegex = regexp.MustCompile(`^(\d+), (\d+)$`)
)

// ParseCommands transforms a command string into a slice of instructions.
// It returns an ErrBadInstruction if the command string contains an invalid instruction
func ParseCommands(commands string) ([]mars.Instruction, error) {
	instructions := make([]mars.Instruction, 0, len(commands))
	for _, command := range commands {
		switch command {
		case 'F', 'L', 'R', 'B':
			instructions = append(instructions, mars.Instruction(command))
		default:
			return nil, mars.ErrBadCommands
		}
	}

	return instructions, nil
}

// ParsePosition parses a string such as `4, 2, EAST` to a matching position
func ParsePosition(pos string) (*position.Position, error) {
	res := positionRegex.FindStringSubmatch(pos)
	if len(res) != 4 { // 3 subgroups plus full match
		return nil, mars.ErrBadPosition
	}

	x, err := strconv.ParseInt(res[1], 10, bits.UintSize)
	if err != nil {
		return nil, mars.ErrBadPosition
	}

	y, err := strconv.ParseInt(res[2], 10, bits.UintSize)
	if err != nil {
		return nil, mars.ErrBadPosition
	}

	var d position.Direction
	switch res[3] {
	case "NORTH":
		d = position.North
	case "SOUTH":
		d = position.South
	case "EAST":
		d = position.East
	case "WEST":
		d = position.West
	default:
		return nil, mars.ErrBadPosition
	}

	return position.NewPosition(position.NewCoordinate(int(x), int(y)), d), nil
}

func ParseCoordinate(coordString string) (coord position.Coordinate, err error) {
	res := coordinateRegex.FindStringSubmatch(coordString)
	if len(res) != 3 { // one for each digit group plus the full match
		return coord, mars.ErrBadCoordinate
	}

	x, err := strconv.ParseInt(res[1], 10, bits.UintSize)
	if err != nil {
		return coord, mars.ErrBadCoordinate
	}

	y, err := strconv.ParseInt(res[2], 10, bits.UintSize)
	if err != nil {
		return coord, mars.ErrBadCoordinate
	}

	coord = position.NewCoordinate(int(x), int(y))
	return coord, nil
}
