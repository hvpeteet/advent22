package cathode_ray

type Instruction interface {
	Step(*CommsDevice)
	Done() bool
}

type CycleInstruction struct {
	RemainingCycles int
}

func (i *CycleInstruction) Done() bool {
	return i.RemainingCycles == 0
}

type NoopInstruction struct {
	CycleInstruction
}

func CreateNoop() *NoopInstruction {
	return &NoopInstruction{CycleInstruction{RemainingCycles: 1}}
}

func (i *NoopInstruction) Step(a *CommsDevice) {
	i.RemainingCycles--
}

type AddInstruction struct {
	CycleInstruction
	amount int
}

func CreateAdd(amount int) *AddInstruction {
	return &AddInstruction{
		CycleInstruction{RemainingCycles: 2},
		amount,
	}
}

func (i *AddInstruction) Step(a *CommsDevice) {
	if i.RemainingCycles == 1 {
		a.X += i.amount
	}
	i.RemainingCycles--
}
