package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/robotlovesyou/mars/position"

	"github.com/robotlovesyou/mars"
	"github.com/robotlovesyou/mars/rover"

	"github.com/robotlovesyou/mars/config"
)

const (
	usage         = `usage: mars [pathToConfig]`
	exampleConfig = `4, 2, EAST
FLFFFRFLB
1, 4
3, 5
7, 4
`
)

// run loads configuration and sends the rover on its journey
func run(confReader *bufio.Reader) (out *position.Position, err error) {
	conf, err := config.Load(confReader)
	if err != nil {
		fmt.Println("Invalid Config. Example:")
		fmt.Println(exampleConfig)
		return out, mars.ErrBadConfig
	}

	rov := rover.New(conf.Start, conf.Map)
	return rov.Execute(conf.Instructions)
}

// prepareOutput prepares the result of the journey for printing
func prepareOutput(pos *position.Position, err error) string {
	if err != nil {
		if errors.Is(mars.ErrBadConfig, err) {
			return ""
		}
		if !errors.Is(mars.ErrStoppedByObstacle, err) {
			return fmt.Sprintf("encountered unexpected error: %v", err)
		}
	}

	posStr := fmt.Sprintf("%v", pos)
	if errors.Is(mars.ErrStoppedByObstacle, err) {
		posStr += " STOPPED"
	}
	return posStr
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(usage)
	}

	fmt.Println(prepareOutput(run(bufio.NewReader(f))))
}
