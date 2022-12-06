package main

import (
	"fmt"

	cc "github.com/hvpeteet/advent22/pkg/4_camp_cleanup"
)

func main() {

	ranges, err := cc.ReadInput("real_input.txt")
	if err != nil {
		panic(err)
	}

	overlaps_count := 0
	for i, range_pair := range ranges {
		if range_pair[0].Overlaps(range_pair[1]) {
			fmt.Printf("Pair %d overlaps: '%v' : %v\n", i, range_pair[0], range_pair[1])
			overlaps_count++
		}
	}
	fmt.Printf("Contains: %d", overlaps_count)
}
