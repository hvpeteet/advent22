package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var exists = struct{}{}

func score_char(char rune) int {
	if unicode.IsUpper(char) {
		return int(char) - int('A') + 27
	} else {
		return int(char) - int('a') + 1
	}
}

func score_group(lines [3]string) int {
	group := find_group(lines)
	fmt.Printf("Group %c\n", group)
	return score_char(group)
}

func find_group(lines [3]string) rune {
	possible := map[rune]struct{}{}
	for _, char := range lines[0] {
		possible[char] = exists
	}
	next_possible := map[rune]struct{}{}
	for _, line := range lines {
		for _, char := range line {
			if _, ok := possible[char]; ok {
				next_possible[char] = exists
			}
		}
		possible = next_possible
		next_possible = map[rune]struct{}{}
	}
	if len(possible) != 1 {
		panic(fmt.Sprintf("more than 1 possible match %v", possible))
	}
	for group := range possible {
		return group
	}
	return '#'
}

func main() {
	file, err := os.Open("real_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	all_scores := []int{}

	scanner := bufio.NewScanner(file)

Readlines:
	for true {
		group_lines := [3]string{}
		for i := range group_lines {
			if !scanner.Scan() {
				break Readlines
			}
			line := scanner.Text()
			group_lines[i] = line
		}
		all_scores = append(all_scores, score_group(group_lines))
	}

	fmt.Printf("%v", all_scores)

	sum_total := 0
	for i := 0; i < len(all_scores); i++ {
		sum_total += all_scores[i]
	}
	fmt.Printf("\n\n%d", sum_total)
}
