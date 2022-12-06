package int_range

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hvpeteet/advent22/pkg/fuck_go"
)

type IntRange struct {
	start int
	end   int
}

func NewIntRange(s string) (IntRange, error) {
	range_str_pieces := strings.Split(s, "-")
	if len(range_str_pieces) != 2 {
		return IntRange{}, fmt.Errorf("range string must only contain 2 ints separated by one '-' was '%s'", s)
	}
	start, end := 0, 0
	var err error
	if start, err = strconv.Atoi(range_str_pieces[0]); err != nil {
		return IntRange{}, err
	}
	if end, err = strconv.Atoi(range_str_pieces[1]); err != nil {
		return IntRange{}, err
	}
	if start > end {
		return IntRange{}, fmt.Errorf("range improperly formatted, first int must be less than second int `%s`", s)
	}
	return IntRange{start: start, end: end}, nil
}

func (r IntRange) Size() int {
	return r.end - r.start
}

func (r1 IntRange) FullyContains(r2 IntRange) bool {
	return r1.start <= r2.start && r1.end >= r2.end
}

func (r1 IntRange) Overlaps(r2 IntRange) bool {
	return r1.start <= r2.end && r1.end >= r2.start
}

func (r1 IntRange) Intersection(r2 IntRange) (IntRange, error) {
	if !r1.Overlaps(r2) {
		return IntRange{}, fmt.Errorf("cannot calculate intersection since no overlap exists")
	}
	return IntRange{
		start: fuck_go.IntMax(r1.start, r2.start),
		end:   fuck_go.IntMin(r1.end, r2.end),
	}, nil
}
