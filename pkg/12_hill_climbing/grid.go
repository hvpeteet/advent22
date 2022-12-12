package hill_climbing

import (
	"container/heap"
	"fmt"

	fg "github.com/hvpeteet/advent22/pkg/fuck_go"
)

type Grid [][]int

func (g Grid) FindAll(height int) []GridPosition {
	ret := []GridPosition{}
	for row_i, row := range g {
		for col_i, val := range row {
			if val == height {
				ret = append(ret, GridPosition{fg.Vector2{X: col_i, Y: row_i}})
			}
		}
	}
	return ret
}

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
	return []GridPosition{}
}
