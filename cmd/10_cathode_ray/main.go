package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	cr "github.com/hvpeteet/advent22/pkg/10_cathode_ray"
)

var fileFlag = flag.String("file", "", "file to parse")

func parseLine(line string) (cr.Instruction, error) {
	parts := strings.Split(line, " ")
	if len(parts) > 2 || len(parts) < 1 {
		return nil, fmt.Errorf("'%s' is not properly formatted", line)
	}
	if parts[0] == "noop" {
		return cr.CreateNoop(), nil
	} else if parts[0] == "addx" {
		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		return cr.CreateAdd(amount), nil
	}
	return nil, fmt.Errorf("'%s' is not properly formatted", line)
}

func main() {
	flag.Parse()

	file, err := os.Open(*fileFlag)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	p1Total := 0
	state := cr.CommsDevice{
		Cycle:              1,
		X:                  1,
		CurrentInstruction: nil,
	}
	for scanner.Scan() {
		line := scanner.Text()

		instruction, err := parseLine(line)
		if err != nil {
			panic(err)
		}

		state.AddInstruction(instruction)
		for state.Busy() {
			if state.Cycle == 20 || (state.Cycle-20)%40 == 0 {
				p1Total += state.SignalStrength()
			}
			state.PrintPixel()
			state.Step()
		}
	}
	fmt.Printf("Total: %d\n", p1Total)
}
