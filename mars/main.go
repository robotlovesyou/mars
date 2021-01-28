package main

import (
	"bufio"
	"fmt"
	"os"

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

func run(confReader *bufio.Reader) (out string, err error) {
	conf, err := config.Load(confReader)
	if err != nil {
		fmt.Println("Invalid Config. Example:")
		fmt.Println(exampleConfig)
		return out, err
	}

	rov := rover.New(conf.Start, conf.Map)
	res, err := rov.Execute(conf.Instructions)
	out = fmt.Sprintf("%v", res)
	return out, err
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

	out, err := run(bufio.NewReader(f))
	if err == mars.ErrStoppedByObstacle {
		out += " STOPPED"
	}
	fmt.Println(out)
}
