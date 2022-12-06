package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	tt "github.com/hvpeteet/advent22/pkg/6_tuning_trouble"
)

var fileFlag = flag.String("file", "", "help message for flag n")

func main() {
	flag.Parse()

	file, err := os.Open(*fileFlag)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Parse stack
	for scanner.Scan() {
		line := scanner.Text()
		if processed, err := tt.FirstUniqueSubstrIndex(line, 4); err != nil {
			panic(err)
		} else {
			fmt.Printf("%d\n", processed)
		}
	}

}
