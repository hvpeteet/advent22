package terminal_state

import (
	"fmt"
	"strconv"
	"strings"
)

type TerminalState struct {
	Root *Node
	Cwd  *Node
}

func (s *TerminalState) PrettyPrint() {
	fmt.Printf("In directory %s, whole tree:\n", s.Cwd.Name)
	s.Root.PrettyPrint()
}

func (s *TerminalState) addFile(pieces []string) error {
	parsed_size, err := strconv.Atoi(pieces[0])
	if err != nil {
		return err
	}
	s.Cwd.Children[pieces[1]] = &Node{
		Parent: s.Cwd,
		Name:   pieces[1],
		Size:   parsed_size,
	}
	s.Cwd.AddSize(parsed_size)
	return nil
}

func (s *TerminalState) addDir(pieces []string) error {
	s.Cwd.Children[pieces[1]] = &Node{
		Name:     pieces[1],
		Parent:   s.Cwd,
		Children: map[string]*Node{},
	}
	return nil
}

func (s *TerminalState) processCmd(pieces []string) error {
	if pieces[1] == "ls" {
		return nil
	} else if pieces[1] == "cd" {
		if len(pieces) != 3 {
			return fmt.Errorf("command '%v' is not a valid 'cd' command", pieces)
		}
		if pieces[2] == "/" {
			s.Cwd = s.Root
		} else if pieces[2] == ".." {
			s.Cwd = s.Cwd.Parent
		} else if child, ok := s.Cwd.Children[pieces[2]]; !ok {
			return fmt.Errorf("%s is not a subdirectory of %s", pieces[2], s.Cwd.Name)
		} else {
			s.Cwd = child
		}
	} else {
		return fmt.Errorf("command '%v' is not a valid command", pieces)
	}
	return nil
}

func (s *TerminalState) Update(line string) error {
	pieces := strings.Split(line, " ")
	if len(pieces) < 2 {
		return fmt.Errorf("line '%s' must contain at least 2 pieces", line)
	}
	var err error
	switch pieces[0] {
	case "$":
		err = s.processCmd(pieces)
	case "dir":
		err = s.addDir(pieces)
	default:
		err = s.addFile(pieces)
	}
	return err
}

func (s *TerminalState) Part1() int {
	m := s.Root.AsMap()
	total := 0
	for _, n := range m {
		if n.Children == nil {
			continue
		}
		if n.Size <= 100000 {
			total += n.Size
		}
	}
	return total
}

func (s *TerminalState) Part2() (string, *Node) {
	m := s.Root.AsMap()
	current_free := 70000000 - s.Root.Size
	target_delete := 30000000 - current_free
	min_matching := s.Root
	min_matching_path := "/"
	for path, n := range m {
		if n.Children == nil {
			continue
		}
		if n.Size >= target_delete && n.Size < min_matching.Size {
			min_matching = n
			min_matching_path = path
		}
	}
	return min_matching_path, min_matching
}
