package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func score_line(line string) int {
	c1 := line[:len(line)/2]
	c2 := line[len(line)/2:]
	mistake := '#'
	for _, char := range c1 {
		if strings.Contains(c2, string(char)) {
			mistake = char
			break
		}
	}
	if unicode.IsUpper(mistake) {
		return int(mistake) - int('A') + 27
	} else {
		return int(mistake) - int('a') + 1
	}

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
