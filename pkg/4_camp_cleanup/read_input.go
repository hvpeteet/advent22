package camp_cleanup

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	ir "github.com/hvpeteet/advent22/pkg/int_range"
)

func parseLine(line string) ([2]ir.IntRange, error) {
	parsed := [2]ir.IntRange{}
	range_pair_string := strings.Split(line, ",")
	if len(range_pair_string) != 2 {
		return parsed, fmt.Errorf("Input is invalid, must be 2 ranges separated by one comma `%s`", line)
	}
	for i, range_string := range range_pair_string {
		if r, err := ir.NewIntRange(range_string); err != nil {
			return [2]ir.IntRange{}, err
		} else {
			parsed[i] = r
		}
	}
	return parsed, nil
}

func ReadInput(path string) ([][2]ir.IntRange, error) {
	file, err := os.Open(path)
	if err != nil {
		return [][2]ir.IntRange{}, err
	}
	defer file.Close()

	ranges := [][2]ir.IntRange{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if parsed, err := parseLine(line); err != nil {
			return [][2]ir.IntRange{}, err
		} else {
			ranges = append(ranges, parsed)
		}
	}
	return ranges, nil
}
