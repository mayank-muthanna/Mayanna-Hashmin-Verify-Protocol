package mpc

import "testing"

func TestShareRecover(t *testing.T) {

	for secret := uint8(0); secret <= 1; secret++ {

		shares, err := ShareSecret(secret)
		if err != nil {
			t.Fatal(err)
		}

		recovered, err := RecoverSecret(shares)
		if err != nil {
			t.Fatal(err)
		}

		if recovered != secret {
			t.Fatalf(
				"expected %d got %d",
				secret,
				recovered,
			)
		}
	}
}

func TestRandomZeroShares(t *testing.T) {

	for i := 0; i < 100; i++ {

		r, err := RandomZeroShares()
		if err != nil {
			t.Fatal(err)
		}

		zero :=
			r[0] ^
				r[1] ^
				r[2]

		if zero != 0 {
			t.Fatal("shares do not xor to zero")
		}
	}
}

func TestANDGate(t *testing.T) {

	inputs := []struct {
		a uint8
		b uint8
	}{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}

	for _, input := range inputs {

		circuit := NewCircuit()

		a := circuit.NewInput()
		b := circuit.NewInput()

		out := circuit.NewAND(a, b)

		circuit.AddOutput(out)

		witness := NewWitness()

		witness.Set(a, input.a)
		witness.Set(b, input.b)

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

		expected := input.a & input.b

		if result != expected {
			t.Fatalf("%d AND %d expected %d got %d", input.a, input.b, expected, result)
		}
	}

}
