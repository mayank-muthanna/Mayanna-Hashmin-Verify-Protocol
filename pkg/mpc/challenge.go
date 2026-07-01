package mpc

import (
	"crypto/sha256"
)

type Challenge uint8

const (
	Open01 Challenge = 0
	Open12 Challenge = 1
	Open20 Challenge = 2
)

func FiatShamir(commitments [3]Commitment) Challenge {

	h := sha256.New()

	for _, commitment := range commitments {
		h.Write(commitment)
	}

	sum := h.Sum(nil)

	return Challenge(
		sum[0] % 3,
	)
}
