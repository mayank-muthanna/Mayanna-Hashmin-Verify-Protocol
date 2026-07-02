package mpc

import "testing"

func TestProveVerify(t *testing.T) {

	circuit := NewCircuit()

	a := circuit.NewInput()
	b := circuit.NewInput()

	out := circuit.NewAND(a, b)

	circuit.AddOutput(out)

	witness := NewWitness()

	witness.Set(a, 1)
	witness.Set(b, 1)

	proof, err := Prove(circuit, witness)

	if err != nil {
		t.Fatal(err)
	}

	if !Verify(circuit, proof) {
		t.Fatal("proof should verify")
	}
}

func TestProveVerifyMnyTimes(t *testing.T) {

	for i := 0; i < 1000; i++ {

		circuit := NewCircuit()

		a := circuit.NewInput()
		b := circuit.NewInput()

		out := circuit.NewAND(a, b)

		circuit.AddOutput(out)

		witness := NewWitness()

		witness.Set(a, 1)
		witness.Set(b, 0)

		proof, err := Prove(circuit, witness)

		if err != nil {
			t.Fatal(err)
		}

		if !Verify(circuit, proof) {
			t.Fatal("proof failed")
		}
	}
}

func TestTamperChallenge(t *testing.T) {

	circuit := NewCircuit()

	a := circuit.NewInput()
	b := circuit.NewInput()

	out := circuit.NewAND(a, b)

	circuit.AddOutput(out)

	witness := NewWitness()

	witness.Set(a, 1)
	witness.Set(b, 1)

	proof, _ := Prove(circuit, witness)

	proof.Rounds[0].Challenge++

	if Verify(circuit, proof) {
		t.Fatal(
			"tampered proof verified",
		)
	}
}

func TestTamperView(t *testing.T) {

	circuit := NewCircuit()

	a := circuit.NewInput()
	b := circuit.NewInput()

	out := circuit.NewAND(a, b)

	circuit.AddOutput(out)

	witness := NewWitness()

	witness.Set(a, 1)
	witness.Set(b, 1)

	proof, _ := Prove(circuit, witness)

	for wire := range proof.Rounds[0].View1.View.Wires {

		proof.Rounds[0].View1.View.Wires[wire] ^= 1

		break
	}

	if Verify(circuit, proof) {
		t.Fatal(
			"tampered proof verified",
		)
	}
}

func TestMissingRounds(t *testing.T) {

	circuit := NewCircuit()

	a := circuit.NewInput()
	b := circuit.NewInput()

	out := circuit.NewAND(a, b)

	circuit.AddOutput(out)

	witness := NewWitness()

	witness.Set(a, 1)
	witness.Set(b, 1)

	proof, _ := Prove(circuit, witness)

	proof.Rounds = proof.Rounds[:10]

	if Verify(circuit, proof) {
		t.Fatal(
			"proof with missing rounds verified",
		)
	}
}

func TestAllTruthTables(t *testing.T) {

	tests := []struct {
		a uint8
		b uint8
	}{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
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

		proof, err := Prove(circuit, witness)

		if err != nil {
			t.Fatal(err)
		}

		if !Verify(circuit, proof) {
			t.Fatalf("failed for %d %d", test.a, test.b)
		}

	}
}
