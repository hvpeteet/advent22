package main

import (
	"flag"
	"fmt"

	cc "github.com/hvpeteet/advent22/pkg/4_camp_cleanup"
)

var fileFlag = flag.String("file", "", "help message for flag n")

func main() {
	flag.Parse()
	ranges, err := cc.ReadInput(*fileFlag)
	if err != nil {
		panic(err)
	}

	contains_count := 0
	for i, range_pair := range ranges {
		if range_pair[0].FullyContains(range_pair[1]) || range_pair[1].FullyContains(range_pair[0]) {
			fmt.Printf("Pair %d is redundant: '%v' : %v\n", i, range_pair[0], range_pair[1])
			contains_count++
		}
	}
	fmt.Printf("Contains: %d", contains_count)
}
