package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func intMax(nums ...int) int {
	max := math.MinInt
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func intMin(nums ...int) int {
	min := math.MaxInt
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

type intRange struct {
	start int
	end   int
}

func newIntRange(s string) (intRange, error) {
	range_str_pieces := strings.Split(s, "-")
	if len(range_str_pieces) != 2 {
		return intRange{}, fmt.Errorf("Range string must only contain 2 ints separated by one '-' was '%s'", s)
	}
	start, end := 0, 0
	var err error
	if start, err = strconv.Atoi(range_str_pieces[0]); err != nil {
		return intRange{}, err
	}
	if end, err = strconv.Atoi(range_str_pieces[1]); err != nil {
		return intRange{}, err
	}
	if start > end {
		return intRange{}, fmt.Errorf("Range improperly formatted, first int must be less than second int `%s`", s)
	}
	return intRange{start: start, end: end}, nil
}

func (self intRange) size() int {
	return self.end - self.start
}

func (self intRange) fullyContains(other intRange) bool {
	return self.start <= other.start && self.end >= other.end
}

func (self intRange) overlaps(other intRange) bool {
	return self.start <= other.end && self.end >= other.start
}

func (self intRange) intersection(other intRange) (intRange, error) {
	if !self.overlaps(other) {
		return intRange{}, fmt.Errorf("Cannot calculate intersection since no overlap exists")
	}
	return intRange{
		start: intMax(self.start, other.start),
		end:   intMin(self.end, other.end),
	}, nil
}

func parseLine(line string) ([2]intRange, error) {
	parsed := [2]intRange{}
	range_pair_string := strings.Split(line, ",")
	if len(range_pair_string) != 2 {
		return parsed, fmt.Errorf("Input is invalid, must be 2 ranges separated by one comma `%s`", line)
	}
	for i, range_string := range range_pair_string {
		if r, err := newIntRange(range_string); err != nil {
			return [2]intRange{}, err
		} else {
			parsed[i] = r
		}
	}
	return parsed, nil
}

func readInput(path string) ([][2]intRange, error) {
	file, err := os.Open(path)
	if err != nil {
		return [][2]intRange{}, err
	}
	defer file.Close()

	ranges := [][2]intRange{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if parsed, err := parseLine(line); err != nil {
			return [][2]intRange{}, err
		} else {
			ranges = append(ranges, parsed)
		}
	}
	return ranges, nil
}

func main() {
	ranges, err := readInput("real_input.txt")
	if err != nil {
		panic(err)
	}

	overlaps_count := 0
	for i, range_pair := range ranges {
		if range_pair[0].overlaps(range_pair[1]) {
			fmt.Printf("Pair %d overlaps: '%v' : %v\n", i, range_pair[0], range_pair[1])
			overlaps_count++
		}
	}
	fmt.Printf("Contains: %d", overlaps_count)
}
