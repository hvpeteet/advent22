package tree_house

func Visible(grid [][]int) [][]bool {
	final_visible := VisibleFromLeft(grid)
	for i := 1; i < 4; i++ {
		partial := RotateNVisible(grid, i)
		for row_i, row := range final_visible {
			for col_i := range row {
				final_visible[row_i][col_i] = final_visible[row_i][col_i] || partial[row_i][col_i]
			}
		}
	}
	return final_visible
}

func VisibleFromLeft(grid [][]int) [][]bool {
	visible := [][]bool{}
	for _, row := range grid {
		row_mask := []bool{}
		max := -1
		for _, h := range row {
			if h > max {
				row_mask = append(row_mask, true)
				max = h
			} else {
				row_mask = append(row_mask, false)
			}
		}
		visible = append(visible, row_mask)
	}
	return visible
}

func RotateNVisible(grid [][]int, n int) [][]bool {
	rotated_grid := grid
	for i := 0; i < n; i++ {
		rotated_grid = RotateGridCW(rotated_grid)
	}
	final_mask := VisibleFromLeft(rotated_grid)
	for i := 0; i < 4-n; i++ {
		final_mask = RotateGridCW(final_mask)
	}
	return final_mask
}
