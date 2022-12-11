package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
}
