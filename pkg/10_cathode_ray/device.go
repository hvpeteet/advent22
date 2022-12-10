package cathode_ray

import "fmt"

type CommsDevice struct {
	Cycle              int
	X                  int
	CurrentInstruction Instruction
}

func (a *CommsDevice) PrintPixel() {
	drawColumn := (a.Cycle - 1) % 40
	if drawColumn == 0 {
		fmt.Printf("\n")
	}
	if a.X == drawColumn || a.X-1 == drawColumn || a.X+1 == drawColumn {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
}

func (a *CommsDevice) Busy() bool {
	return !a.CurrentInstruction.Done()
}

func (a *CommsDevice) SignalStrength() int {
	return a.X * a.Cycle
}

func (a *CommsDevice) AddInstruction(i Instruction) error {
	if a.CurrentInstruction == nil || a.CurrentInstruction.Done() {
		a.CurrentInstruction = i
		return nil
	}
	return fmt.Errorf("we don't support queuing instructions, current: %+v requested: %v", a.CurrentInstruction, i)

}

func (a *CommsDevice) Step() error {
	if a.CurrentInstruction == nil {
		return fmt.Errorf("can't step with a nil instruction")
	}
	a.CurrentInstruction.Step(a)
	a.Cycle++
	return nil
}
