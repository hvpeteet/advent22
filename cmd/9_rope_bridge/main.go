package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	rb "github.com/hvpeteet/advent22/pkg/9_rope_bridge"
)

var fileFlag = flag.String("file", "", "file to parse")
var knotsFlag = flag.Int("knots", 2, "number of knots")

var moves = map[string]rb.Vector2{
	"U": rb.Up,
	"D": rb.Down,
	"L": rb.Left,
	"R": rb.Right,
}

var Exists = struct{}{}

func parseLine(line string) (rb.Vector2, int, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return rb.Vector2{}, 0, fmt.Errorf("'%s' is not properly formatted", line)
	}
	move := parts[0]
	distance, err := strconv.Atoi(parts[1])
	if err != nil {
		return rb.Vector2{}, 0, fmt.Errorf("parse error on line %s, %e", line, err)
	}
	return moves[move], distance, nil
}

func main() {
	flag.Parse()

	file, err := os.Open(*fileFlag)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tail_locations := map[rb.Vector2]struct{}{}
	rope := rb.NewRope(*knotsFlag)
	fmt.Printf("%+v\n\n", rope)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
		move, distance, err := parseLine(line)
		if err != nil {
			panic(err)
		}

		for i := 0; i < distance; i++ {
			rope.Move(move)
			tail_locations[rope.Knots[len(rope.Knots)-1]] = Exists
			fmt.Printf("%+v\n\n", rope)
		}
	}
	fmt.Printf("%d locations visited, what well traveled rope... knot\n", len(tail_locations))
}
