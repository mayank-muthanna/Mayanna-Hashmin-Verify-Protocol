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

	x := [3]uint8{}
	y := [3]uint8{}

	for i := 0; i < 3; i++ {
		x[i] = views[i].Wires[gate.InputA]
		y[i] = views[i].Wires[gate.InputB]
	}

	r, err := RandomZeroShares()
	if err != nil {
		return err
	}

	z0 := (x[0] & y[0]) ^ (x[0] & y[1]) ^ (x[1] & y[0]) ^ r[0] ^ r[2]

	z1 := (x[1] & y[1]) ^ (x[1] & y[2]) ^ (x[2] & y[1]) ^ r[1] ^ r[0]

	z2 := (x[2] & y[2]) ^ (x[2] & y[0]) ^ (x[0] & y[2]) ^ r[2] ^ r[1]

	views[0].Wires[gate.Output] = z0
	views[1].Wires[gate.Output] = z1
	views[2].Wires[gate.Output] = z2

	views[0].Randomness = append(
		views[0].Randomness,
		r[0],
	)

	views[1].Randomness = append(
		views[1].Randomness,
		r[1],
	)

	views[2].Randomness = append(
		views[2].Randomness,
		r[2],
	)

	return nil
}
