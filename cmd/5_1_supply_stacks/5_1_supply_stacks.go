package main

import (
	"flag"
	"fmt"

	ss "github.com/hvpeteet/advent22/pkg/5_supply_stacks"
)

var fileFlag = flag.String("file", "", "help message for flag n")

func main() {
	flag.Parse()
	state, actions, err := ss.ReadInput(*fileFlag, false)
	if err != nil {
		panic(err)
	}

	for _, action := range *actions {
		fmt.Printf("State: %+v\n", state)
		fmt.Printf("Action: %+v\n", action)
		action.Execute(state)
	}

	for _, stack := range (*state).Stacks {
		if len(stack) > 0 {
			fmt.Printf("%c", stack[len(stack)-1])
		} else {
			fmt.Printf(" ")
		}
	}
	fmt.Print("\n")
}
