package intcomp

// Computer - advent of code Int Comp
type Computer struct {
	Program            []int
	instructionPointer int
	halt               bool
	Input              int
	output             int
}

func (c *Computer) getCurrentOpCode() int {
	rawOpCode := c.Program[c.instructionPointer]
	return rawOpCode % 100
}

func (c *Computer) getRawParameter(index int) int {
	return c.Program[c.instructionPointer+index+1]
}

func (c *Computer) getParameter(index int) int {
	value := c.getRawParameter(index)
	rawOpCode := c.Program[c.instructionPointer]
	mode := 0

	switch index {
	case 0:
		mode = (rawOpCode % 1000) / 100
	case 1:
		mode = (rawOpCode % 10000) / 1000
	}

	if mode == 0 {
		return c.Program[value]
	}

	return value
}

func (c *Computer) logic() {
	opCode := c.getCurrentOpCode()

	switch opCode {
	case 1:
		c.Program[c.getRawParameter(2)] = c.getParameter(0) + c.getParameter(1)
	case 2:
		c.Program[c.getRawParameter(2)] = c.getParameter(0) * c.getParameter(1)
	case 3:
		c.Program[c.getRawParameter(0)] = c.Input
	case 4:
		c.output = c.getParameter(0)
	case 5:
	case 6:
	case 7:
		c.Program[c.getRawParameter(2)] = booleanInt(c.getParameter(0) < c.getParameter(1))
	case 8:
		c.Program[c.getRawParameter(2)] = booleanInt(c.getParameter(0) == c.getParameter(1))
	case 99:
		fallthrough
	default:
		c.halt = true
	}
}

func (c *Computer) movePointer() {
	opCode := c.getCurrentOpCode()

	switch opCode {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 7:
		fallthrough
	case 8:
		c.instructionPointer += 4
	case 3:
		fallthrough
	case 4:
		c.instructionPointer += 2
	case 5:
		if c.getParameter(0) != 0 {
			c.instructionPointer = c.getParameter(1)
		} else {
			c.instructionPointer += 3
		}
	case 6:
		if c.getParameter(0) == 0 {
			c.instructionPointer = c.getParameter(1)
		} else {
			c.instructionPointer += 3
		}
	}
}

// RunProgram - run program on Int Comp
func (c *Computer) RunProgram() int {
	c.halt = false
	c.instructionPointer = 0

	for !c.halt {
		c.logic()
		c.movePointer()
	}

	return c.output
}

func booleanInt(condition bool) int {
	if condition {
		return 1
	}

	return 0
}
