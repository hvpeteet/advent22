package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("real_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	total_calories := []int{0}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			total_calories = append(total_calories, 0)
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		total_calories[len(total_calories)-1] += calories

	}
	sort.Ints(total_calories)
	fmt.Printf("%v", total_calories)

	sum_total := 0
	for i := 1; i <= 3; i++ {
		sum_total += total_calories[len(total_calories)-i]
	}
	fmt.Printf("\n\n%d", sum_total)
}
