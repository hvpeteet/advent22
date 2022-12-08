package tree_house

func ViewDistLeftRow(row []int) []int {
	dists := []int{}
	// If a tree at height x is found what is the furthest index left it can see.
	view_index_per_height := [10]int{0}
	for col_i, h := range row {
		dists = append(dists, col_i-view_index_per_height[h])
		for i := 0; i <= h; i++ {
			view_index_per_height[i] = col_i
		}
	}
	return dists
}

func ViewDistLeft(grid [][]int) [][]int {
	dists := [][]int{}
	for _, row := range grid {
		view_distances := ViewDistLeftRow(row)
		dists = append(dists, view_distances)
	}
	return dists
}

func ViewDistFromRotation(grid [][]int, n int) [][]int {
	rotated_grid := grid
	for i := 0; i < n; i++ {
		rotated_grid = RotateGridCW(rotated_grid)
	}
	final_view_dist := ViewDistLeft(rotated_grid)
	for i := 0; i < 4-n; i++ {
		final_view_dist = RotateGridCW(final_view_dist)
	}
	return final_view_dist
}

func ScenicScores(grid [][]int) [][]int {
	final_scores := ViewDistFromRotation(grid, 0)
	for i := 1; i < 4; i++ {
		partial := ViewDistFromRotation(grid, i)
		for row_i, row := range final_scores {
			for col_i := range row {
				final_scores[row_i][col_i] = final_scores[row_i][col_i] * partial[row_i][col_i]
			}
		}
	}
	return final_scores
}
