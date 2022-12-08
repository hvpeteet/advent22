package tree_house

import (
	"bufio"
	"os"
	"strconv"
)

func ParseInput(f string) ([][]int, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	heights := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, c := range line {
			val, err := strconv.Atoi(string(c))
			if err != nil {
				return nil, err
			}
			row = append(row, val)
		}
		heights = append(heights, row)
	}
	return heights, nil
}
