package hill_climbing

import (
	fg "github.com/hvpeteet/advent22/pkg/fuck_go"
)

type GridPosition struct {
	fg.Vector2
}

var neighborOffsets = []fg.Vector2{fg.Up, fg.Down, fg.Left, fg.Right}

func (g GridPosition) Neighbors(grid Grid) []GridPosition {
	height := grid[g.Y][g.X]
	neighbors := []GridPosition{}
	for _, offset := range neighborOffsets {
		neighborCoords := g.Add(offset)
		if !grid.InBounds(neighborCoords) {
			continue
		}
		diff := grid.Height(neighborCoords) - height
		if grid.Height(neighborCoords) > height && diff > 1 {
			continue
		}
		neighbors = append(neighbors, GridPosition{neighborCoords})
	}
	return neighbors
}

func Lines2Grid(lines []string) (Grid, GridPosition, GridPosition) {
	g := Grid{}
	var start, end GridPosition
	for row_i, line := range lines {
		row := []int{}
		for col_i, c := range line {
			if c == 'S' {
				start = GridPosition{fg.Vector2{X: col_i, Y: row_i}}
				c = 'a'
			}
			if c == 'E' {
				end = GridPosition{fg.Vector2{X: col_i, Y: row_i}}
				c = 'z'
			}
			row = append(row, int(c)-int('a'))
		}
		g = append(g, row)
	}
	return g, start, end
}
