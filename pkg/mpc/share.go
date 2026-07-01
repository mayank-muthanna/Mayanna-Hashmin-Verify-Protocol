package mpc

import (
	"crypto/rand"
	"fmt"
)

type Share struct {
	Value uint8
}

func ShareSecret(secret uint8) ([3]Share, error) {

	var shares [3]Share

	random := make([]byte, 2)

	_, err := rand.Read(random)

	if err != nil {
		return shares, err
	}

	shares[0].Value = random[0] & 1
	shares[1].Value = random[1] & 1
	shares[2].Value = secret ^ shares[0].Value ^ shares[1].Value

	return shares, nil
}

func RecoverSecret(shares [3]Share) (uint8, error) {

	secret := shares[0].Value ^ shares[1].Value ^ shares[2].Value

	if secret > 1 {
		return 0, fmt.Errorf("Invalid Secret")
	}

	return secret, nil

}
