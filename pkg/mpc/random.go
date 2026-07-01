package mpc

import "crypto/rand"

func RandomBit() (uint8, error) {

	b := make([]byte, 1)

	_, err := rand.Read(b)

	if err != nil {
		return 0, err
	}

	return b[0] & 1, nil
}

func RandomZeroShares() ([3]uint8, error) {

	var shares [3]uint8

	r0, err := RandomBit()
	if err != nil {
		return shares, err
	}

	r1, err := RandomBit()
	if err != nil {
		return shares, err
	}

	r2 := r0 ^ r1

	shares[0] = r0
	shares[1] = r1
	shares[2] = r2

	return shares, nil
}
