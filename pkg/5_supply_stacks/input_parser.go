package supply_stacks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseCommand(cmd string, mass_move bool) (Command, error) {
	parts := strings.Split(cmd, " ")
	if len(parts) != 6 {
		return nil, fmt.Errorf("malformed command '%s'", cmd)
	}
	var parsed_ints []int
	for _, i := range []int{1, 3, 5} {
		if parsed, err := strconv.Atoi(parts[i]); err != nil {
			return nil, err
		} else {
			parsed_ints = append(parsed_ints, parsed)
		}
	}
	if mass_move {
		return &MassMoveCommand{
			Quantity: parsed_ints[0],
			From:     parsed_ints[1],
			To:       parsed_ints[2],
		}, nil
	}
	return &MoveCommand{
		Quantity: parsed_ints[0],
		From:     parsed_ints[1],
		To:       parsed_ints[2],
	}, nil
}

func AccumulateState(state *State, line []rune) {
	stack := 0
	for i := 0; i+2 <= len(line)-1; i += 4 {
		stack++
		if len(state.Stacks) < stack {
			state.Stacks = append(state.Stacks, []rune{})
		}
		if line[i] != '[' {
			continue
		}
		state.Stacks[stack-1] = append([]rune{line[i+1]}, state.Stacks[stack-1]...)
	}
}

func ReadInput(path string, mass_move bool) (*State, *[]Command, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var state State
	var commands []Command

	scanner := bufio.NewScanner(file)
	// Parse stack
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		AccumulateState(&state, []rune(line))
	}
	// Parse commands
	for scanner.Scan() {
		line := scanner.Text()
		if cmd, err := ParseCommand(line, mass_move); err != nil {
			panic(err)
		} else {
			commands = append(commands, cmd)
		}

	}
	return &state, &commands, nil
}
