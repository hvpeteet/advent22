package tree_house

func GridMax(grid [][]int) int {
	max := grid[0][0]
	for _, r := range grid {
		for _, v := range r {
			if v > max {
				max = v
			}
		}
	}
	return max
}
