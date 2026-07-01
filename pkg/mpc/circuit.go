package mpc

import "fmt"

type WireID uint32

const (
	FalseWire WireID = 0
	TrueWire  WireID = 1
)

type GateType string

const (
	AND GateType = "AND"
	XOR GateType = "XOR"
	NOT GateType = "NOT"
)

type Gate struct {
	GateType GateType

	InputA WireID
	InputB WireID

	Output WireID
}

type Circuit struct {
	Inputs  []WireID
	Outputs []WireID

	Gates []Gate

	nextWire WireID
}

func NewCircuit() *Circuit {

	return &Circuit{
		nextWire: 2,
	}

}

func (c *Circuit) NewInput() WireID {

	wire := c.nextWire
	c.nextWire++
	c.Inputs = append(c.Inputs, wire)
	return wire
}

func (c *Circuit) NewAND(a, b WireID) WireID {

	output := c.nextWire

	c.nextWire++

	c.Gates = append(c.Gates, Gate{

		GateType: AND,

		InputA: a,
		InputB: b,

		Output: output,
	})

	return output
}

func (c *Circuit) NewXOR(a, b WireID) WireID {

	output := c.nextWire

	c.nextWire++

	c.Gates = append(c.Gates, Gate{

		GateType: XOR,

		InputA: a,
		InputB: b,

		Output: output,
	})

	return output
}

func (c *Circuit) NewNOT(a WireID) WireID {

	output := c.nextWire

	c.nextWire++

	c.Gates = append(c.Gates, Gate{

		GateType: NOT,

		InputA: a,

		Output: output,
	})

	return output
}

func (c *Circuit) AddOutput(wire WireID) {

	c.Outputs = append(c.Outputs, wire)

}

func (c *Circuit) TotalWires() int {

	return int(c.nextWire)

}

func (c *Circuit) Validate() error {

	for _, gate := range c.Gates {

		if gate.Output == FalseWire || gate.Output == TrueWire {
			return fmt.Errorf("cannot write to constant wires")
		}

		if gate.InputA >= c.nextWire {
			return fmt.Errorf("invalid input wire")
		}

		if gate.GateType != NOT {

			if gate.InputB >= c.nextWire {
				return fmt.Errorf("invalid input wire")
			}
		}

		if gate.Output >= c.nextWire {
			return fmt.Errorf("invalid output wire")
		}
	}

	return nil
}
