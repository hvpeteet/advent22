package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	ns "github.com/hvpeteet/advent22/pkg/7_no_space"
)

var fileFlag = flag.String("file", "", "help message for flag n")

func main() {
	flag.Parse()

	file, err := os.Open(*fileFlag)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	root := &ns.Node{
		Name:     "/",
		Children: map[string]*ns.Node{},
	}
	state := ns.TerminalState{
		Root: root,
		Cwd:  root,
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err := state.Update(line); err != nil {
			panic(err)
		}
	}
	state.PrettyPrint()
	fmt.Printf("Part1 Count: %d\n", state.Part1())
	path, node := state.Part2()
	fmt.Printf("Part2: %s :%d\n", path, node.Size)
}
