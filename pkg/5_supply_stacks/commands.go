package supply_stacks

type Command interface {
	Execute(state *State)
}

type MoveCommand struct {
	Quantity int
	From     int
	To       int
}

func (cmd MoveCommand) Execute(state *State) {
	for i := 0; i < cmd.Quantity; i++ {
		from := &state.Stacks[cmd.From-1]
		to := &state.Stacks[cmd.To-1]
		if len(*from) == 0 {
			break
		}
		state.Stacks[cmd.To-1] = append(*to, (*from)[len(*from)-1])
		state.Stacks[cmd.From-1] = (*from)[:len(*from)-1]
	}
}

type MassMoveCommand struct {
	Quantity int
	From     int
	To       int
}

func (cmd MassMoveCommand) Execute(state *State) {
	from := state.Stacks[cmd.From-1]
	to := state.Stacks[cmd.To-1]
	state.Stacks[cmd.To-1] = append(to, from[len(from)-cmd.Quantity:]...)
	state.Stacks[cmd.From-1] = state.Stacks[cmd.From-1][:len(from)-cmd.Quantity]
}
