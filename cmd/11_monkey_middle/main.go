package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/big"
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
	for _, x := range parse.AllInts(strings.Replace(lines[1], ",", "", -1)) {
		m.Items = append(m.Items, big.NewInt(int64(x)))
	}
	m.Operation = lines[2]
	m.TestDivsibleBy = big.NewInt(int64(parse.AllInts(lines[3])[0])) // TODO: Check for empty array
	m.TrueTarget = parse.AllInts(lines[4])[0]                        // TODO: Check for empty array
	m.FalseTarget = parse.AllInts(lines[5])[0]                       // TODO: Check for empty array
	return &m, nil

}

func monkeyBuisness(monkeys []*mky.Monkey) int64 {
	itemsInspected := make([]int64, len(monkeys))
	// fmt.Print("------\n\n")
	for round := 0; round < 10000; round++ {
		for monkey_i, monkey := range monkeys {
			for _, worry := range monkey.Items {
				if adjustedWorry, err := monkey.ApplyOperation(worry); err != nil {
					panic(err)
				} else {
					// adjustedWorry.Div(adjustedWorry, big.NewInt(3))
					targetMonkey := monkey.FalseTarget
					if big.NewInt(0).Mod(adjustedWorry, monkey.TestDivsibleBy).CmpAbs(big.NewInt(0)) == 0 {
						targetMonkey = monkey.TrueTarget
					}
					// fmt.Printf("Throwing item with worry %d to monkey %d\n", adjustedWorry, targetMonkey)
					monkeys[targetMonkey].Items = append(monkeys[targetMonkey].Items, adjustedWorry)
				}
				itemsInspected[monkey_i]++
			}
			monkey.Items = []*big.Int{}
		}
		// for _, m := range monkeys {
		// 	fmt.Printf("%+v\n", m)
		// }
		// fmt.Print("------\n\n")
		if round%500 == 0 {
			fmt.Printf("%d\n", round)
			for _, m := range monkeys {
				fmt.Printf("%+v\n", m)
			}
		}
	}
	sort.Slice(itemsInspected, func(i, j int) bool { return itemsInspected[i] < itemsInspected[j] })
	return itemsInspected[len(itemsInspected)-1] * itemsInspected[len(itemsInspected)-2]
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
