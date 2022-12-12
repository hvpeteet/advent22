package main

import (
	"bufio"
	"container/heap"
	"flag"
	"fmt"
	"os"

	fg "github.com/hvpeteet/advent22/pkg/fuck_go"
)

var fileFlag = flag.String("file", "", "file to parse")

func readAllLines() ([]string, error) {
	flag.Parse()

	file, err := os.Open(*fileFlag)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines, nil
}

type Grid [][]int

func (g Grid) PrettyPrint() {
	for _, row := range g {
		for _, val := range row {
			fmt.Printf("%02d ", val)
		}
		fmt.Println()
	}
	fmt.Println()
}

// Top left is alway 0,0
func (g Grid) Height(pos fg.Vector2) int {
	return g[pos.Y][pos.X]
}

func (g Grid) InBounds(pos fg.Vector2) bool {
	return pos.Y < len(g) && pos.X < len(g[0]) && pos.Y >= 0 && pos.X >= 0
}

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

func lines2Grid(lines []string) (Grid, GridPosition, GridPosition) {
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

type SearchItem struct {
	Path []GridPosition
}

var exists interface{}

func (g Grid) Search(start, end GridPosition) []GridPosition {
	pq := make(fg.PriorityQueue[SearchItem], 0)
	heap.Init(&pq) // Likely not needed since it is emtpy / size 1
	pq.PushT(0, SearchItem{Path: []GridPosition{start}})
	visited := map[GridPosition]any{}
	for len(pq) > 0 {
		item := pq.PopT()
		startNode := item.Path[len(item.Path)-1]
		if _, ok := visited[startNode]; ok {
			continue
		}
		for _, n := range startNode.Neighbors(g) {
			if n == end {
				return fg.NonMutateAppend(item.Path, n)
			}
			priority := n.Subtract(fg.Vector2{X: end.X, Y: end.Y}).Manhattan() + len(item.Path) + 1
			pq.PushT(priority, SearchItem{Path: fg.NonMutateAppend(item.Path, n)})
		}
		visited[startNode] = exists
	}
	fmt.Printf("visited: %v\n", visited)
	return []GridPosition{}
}

func main() {
	// Read input in
	lines, err := readAllLines()
	if err != nil {
		panic(err)
	}
	grid, start, end := lines2Grid(lines)
	grid.PrettyPrint()
	// A* search (traditional manhattan heuristic)
	path := grid.Search(start, end)
	// Return len(path)
	for _, n := range path {
		fmt.Printf("%+v\n", n)
	}
	fmt.Printf("len: %d\n", len(path)-1)
}
