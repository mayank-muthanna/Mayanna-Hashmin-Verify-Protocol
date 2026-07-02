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
