package mpc

import "fmt"

func EvaluateCircuit(
	circuit *Circuit,
	witness *Witness,
) ([]*View, error) {

	views := []*View{
		NewView(0),
		NewView(1),
		NewView(2),
	}

	// 1. Share every input wire
	for _, wire := range circuit.Inputs {

		secret := witness.Get(wire)
		shares, err := ShareSecret(secret)

		if err != nil {
			return nil, err
		}

		for i := 0; i < 3; i++ {
			views[i].Wires[wire] = shares[i].Value
		}
	}

	// 2. Evaluate every gate
	for _, gate := range circuit.Gates {

		switch gate.GateType {

		case XOR:

			evaluateXOR(gate, views)

		case NOT:

			evaluateNOT(gate, views)

		case AND:

			err := evaluateAND(gate, views)

			if err != nil {
				return nil, err
			}

		default:
			return nil, fmt.Errorf("unknown gate")
		}
	}

	return views, nil
}

func evaluateXOR(gate Gate, views []*View) {
	for _, view := range views {

		view.Wires[gate.Output] = view.Wires[gate.InputA] ^ view.Wires[gate.InputB]
	}
}

func evaluateNOT(gate Gate, views []*View) {

	for _, view := range views {

		view.Wires[gate.Output] = view.Wires[gate.InputA] ^ 1
	}
}

func evaluateAND(gate Gate, views []*View) error {

	for _, view := range views {

		a := view.Wires[gate.InputA]

		b := view.Wires[gate.InputB]

		view.Wires[gate.Output] = a & b
	}

	return nil
}
