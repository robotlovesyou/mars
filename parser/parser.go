package parser

import mars "github.com/robotlovesyou/mars"

func Parse(commands string) ([]mars.Instruction, error) {
	instructions := make([]mars.Instruction, 0, len(commands))
	for _, command := range commands {
		switch (command) {
		case 'F', 'L', 'R', 'B':
			instructions = append(instructions, mars.Instruction(command))
		default:
			return nil, mars.ErrBadCommands
		}
	}

	return instructions, nil
}