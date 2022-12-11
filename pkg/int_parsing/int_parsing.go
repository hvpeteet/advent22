package int_parsing

import (
	"strconv"
	"strings"
)

func AllInts(line string) []int {
	all := []int{}
	for _, p := range strings.Split(line, " ") {
		if i, err := strconv.Atoi(p); err != nil {
			continue
		} else {
			all = append(all, i)
		}
	}
	return all
}
