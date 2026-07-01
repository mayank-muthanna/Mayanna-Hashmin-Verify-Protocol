package mpc

import "crypto/rand"

type Challenge uint8

const (
	Open01 Challenge = 0
	Open12 Challenge = 1
	Open20 Challenge = 2
)

func RandomChallenge() (Challenge, error) {

	b := make([]byte, 1)

	_, err := rand.Read(b)
	if err != nil {
		return 0, nil
	}

	return Challenge(b[0] % 3), nil
}
