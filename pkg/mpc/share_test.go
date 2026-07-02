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
