package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	mky "github.com/hvpeteet/advent22/pkg/11_monkey_middle"
	parse "github.com/hvpeteet/advent22/pkg/int_parsing"
)

var fileFlag = flag.String("file", "", "file to parse")

func parseMonkeyNote(lines []string) (*mky.Monkey, error) {
	if len(lines) != 6 {
		return nil, fmt.Errorf("'%v' is not properly formatted", lines)
	}
	m := mky.Monkey{}
	m.Items = parse.AllInts(strings.Replace(lines[1], ",", "", -1))
	m.Operation = lines[2]
	m.TestDivsibleBy = parse.AllInts(lines[3])[0] // TODO: Check for empty array
	m.TrueTarget = parse.AllInts(lines[4])[0]     // TODO: Check for empty array
	m.FalseTarget = parse.AllInts(lines[5])[0]    // TODO: Check for empty array
	return &m, nil

}

func monkeyBuisness(monkeys []*mky.Monkey) int {
	itemsInspected := make([]int, len(monkeys))
	fmt.Print("------\n\n")
	for round := 0; round < 20; round++ {
		for monkey_i, monkey := range monkeys {
			for _, worry := range monkey.Items {
				if adjustedWorry, err := monkey.ApplyOperation(worry); err != nil {
					panic(err)
				} else {
					// adjustedWorry = adjustedWorry / 3
					targetMonkey := monkey.FalseTarget
					if adjustedWorry%monkey.TestDivsibleBy == 0 {
						targetMonkey = monkey.TrueTarget
					}
					fmt.Printf("Throwing item with worry %d to monkey %d\n", adjustedWorry, targetMonkey)
					monkeys[targetMonkey].Items = append(monkeys[targetMonkey].Items, adjustedWorry)
				}
				itemsInspected[monkey_i]++
			}
			monkey.Items = []int{}
		}
		for _, m := range monkeys {
			fmt.Printf("%+v\n", m)
		}
		fmt.Print("------\n\n")
	}
	sort.Sort(sort.Reverse(sort.IntSlice(itemsInspected)))
	return itemsInspected[0] * itemsInspected[1]
}

func main() {
	flag.Parse()

	file, err := os.Open(*fileFlag)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	monkeys := []*mky.Monkey{}
	for i := 0; i < len(lines); i += 6 {
		if m, err := parseMonkeyNote(lines[i : i+6]); err != nil {
			panic(err)
		} else {
			monkeys = append(monkeys, m)
		}
	}
	for _, m := range monkeys {
		fmt.Printf("%+v\n", m)
	}
	fmt.Printf("Part1: %d\n", monkeyBuisness(monkeys))
}
