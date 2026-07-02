package mpc

import "testing"

func TestXORGate(t *testing.T) {

	tests := []struct {
		a        uint8
		b        uint8
		expected uint8
	}{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
	}

	for _, test := range tests {

		circuit := NewCircuit()

		a := circuit.NewInput()
		b := circuit.NewInput()

		out := circuit.NewXOR(a, b)

		circuit.AddOutput(out)

		witness := NewWitness()

		witness.Set(a, test.a)
		witness.Set(b, test.b)

		views, err := EvaluateCircuit(circuit, witness)

		if err != nil {
			t.Fatal(err)
		}

		var shares [3]Share

		for i := 0; i < 3; i++ {
			shares[i] = Share{
				Value: views[i].Wires[out],
			}
		}

		result, err := RecoverSecret(shares)
		if err != nil {
			t.Fatal(err)
		}

		if result != test.expected {
			t.Fatalf("%d XOR %d expected %d got %d", test.a, test.b, test.expected, result)
		}
	}
}

func TestANDGate(t *testing.T) {

	tests := []struct {
		a        uint8
		b        uint8
		expected uint8
	}{
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
		{1, 1, 1},
	}

	for _, test := range tests {

		circuit := NewCircuit()

		a := circuit.NewInput()
		b := circuit.NewInput()

		out := circuit.NewAND(a, b)

		circuit.AddOutput(out)

		witness := NewWitness()

		witness.Set(a, test.a)
		witness.Set(b, test.b)

		views, err := EvaluateCircuit(
			circuit,
			witness,
		)

		if err != nil {
			t.Fatal(err)
		}

		var shares [3]Share

		for i := 0; i < 3; i++ {
			shares[i] = Share{
				Value: views[i].Wires[out],
			}
		}

		result, err := RecoverSecret(shares)
		if err != nil {
			t.Fatal(err)
		}

		if result != test.expected {
			t.Fatalf(
				"%d AND %d expected %d got %d",
				test.a,
				test.b,
				test.expected,
				result,
			)
		}
	}
}

func TestNOTGate(t *testing.T) {

	tests := []struct {
		input    uint8
		expected uint8
	}{
		{0, 1},
		{1, 0},
	}

	for _, test := range tests {

		circuit := NewCircuit()

		in := circuit.NewInput()

		out := circuit.NewNOT(in)

		circuit.AddOutput(out)

		witness := NewWitness()

		witness.Set(in, test.input)

		views, err := EvaluateCircuit(circuit, witness)

		if err != nil {
			t.Fatal(err)
		}

		var shares [3]Share

		for i := 0; i < 3; i++ {
			shares[i] = Share{
				Value: views[i].Wires[out],
			}
		}

		result, err := RecoverSecret(shares)
		if err != nil {
			t.Fatal(err)
		}

		if result != test.expected {
			t.Fatalf("NOT %d expected %d got %d", test.input, test.expected, result)
		}

	}
}
