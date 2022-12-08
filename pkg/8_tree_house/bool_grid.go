package tree_house

func CountTrue(visible [][]bool) int {
	total := 0
	for _, r := range visible {
		for _, v := range r {
			if v {
				total++
			}
		}
	}
	return total
}
