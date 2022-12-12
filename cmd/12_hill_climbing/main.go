package main

import (
	"flag"
	"fmt"
	"math"

	hc "github.com/hvpeteet/advent22/pkg/12_hill_climbing"
	fg "github.com/hvpeteet/advent22/pkg/fuck_go"
)

var fileFlag = flag.String("file", "", "file to parse")
var part1Flag = flag.Bool("part1", true, "part1")

func main() {
	flag.Parse()
	// Read input in
	lines, err := fg.ReadAllLines(*fileFlag)
	if err != nil {
		panic(err)
	}
	grid, start, end := hc.Lines2Grid(lines)
	// A* search (traditional manhattan heuristic)
	if *part1Flag {
		path := grid.Search(start, end)
		fmt.Printf("len: %d\n", len(path)-1)
	} else {
		starts := grid.FindAll(0)
		min := math.MaxInt
		for _, s := range starts {
			path := grid.Search(s, end)
			if len(path) > 0 {
				min = fg.IntMin(min, len(path))
			}
		}
		fmt.Printf("len: %d\n", min-1)
	}

}
