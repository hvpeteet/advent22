package main

import (
	"flag"
	"fmt"

	th "github.com/hvpeteet/advent22/pkg/8_tree_house"
)

var fileFlag = flag.String("file", "", "help message for flag n")

func main() {
	flag.Parse()

	heights, err := th.ParseInput(*fileFlag)
	if err != nil {
		panic(err)
	}
	scores := th.ScenicScores(heights)
	fmt.Printf("Max: %d\n", th.GridMax(scores))

}
