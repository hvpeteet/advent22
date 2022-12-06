package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var substitution_codex = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
}

var beats = map[string]string{
	"Rock":     "Scissors",
	"Paper":    "Rock",
	"Scissors": "Paper",
}

var loses = map[string]string{
	"Scissors": "Rock",
	"Rock":     "Paper",
	"Paper":    "Scissors",
}

var outcome_cipher = map[string]string{
	"X": "Lose",
	"Y": "Tie",
	"Z": "Win",
}

var points_per_selection = map[string]int{
	"Rock":     1,
	"Paper":    2,
	"Scissors": 3,
}

func score_line(line string) int {
	score := 0
	selections := strings.Split(line, " ")
	if len(selections) != 2 {
		panic(fmt.Sprintf("line %s was not correctly formatted", line))
	}
	other_play := substitution_codex[selections[0]]
	outcome := outcome_cipher[selections[1]]
	my_play := "Undecided"
	switch outcome {
	case "Win":
		my_play = loses[other_play]
	case "Tie":
		my_play = other_play
	case "Lose":
		my_play = beats[other_play]
	}
	// Score from selection
	score += points_per_selection[my_play]
	// Competitive score
	if other_play == my_play {
		score += 3
		return score
	} else if beats[my_play] == other_play {
		score += 6
	}
	return score
}

func main() {
	file, err := os.Open("real_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	all_scores := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		all_scores = append(all_scores, score_line(line))
	}
	fmt.Printf("%v", all_scores)

	sum_total := 0
	for i := 0; i < len(all_scores); i++ {
		sum_total += all_scores[i]
	}
	fmt.Printf("\n\n%d", sum_total)
}
