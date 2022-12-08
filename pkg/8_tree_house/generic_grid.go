package tree_house

import "fmt"

func RotateGridCW[T interface{}](grid [][]T) [][]T {
	rotated := [][]T{}
	for final_x := 0; final_x < len(grid[0]); final_x++ {
		final_row := []T{}
		for final_y := 0; final_y < len(grid); final_y++ {
			final_row = append(final_row, grid[len(grid)-final_y-1][final_x])
		}
		rotated = append(rotated, final_row)
	}
	return rotated
}

func PrintGrid[T interface{}](grid [][]T) {
	for _, r := range grid {
		fmt.Printf("%v\n", r)
	}
}
