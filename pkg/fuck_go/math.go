package fuck_go

import "math"

func IntMax(nums ...int) int {
	max := math.MinInt
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func IntMin(nums ...int) int {
	min := math.MaxInt
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IntSign(x int) int {
	if x >= 0 {
		return 1
	}
	return -1
}
