package hm

import "crypto/rand"

type BitCommitment struct {
	Index      int
	Commitment []byte
}

type BitOpening struct {
	Index int
	Bit   uint8
	Salt  []byte
}

const SaltSize = 32

func RandomSalt() ([]byte, error) {
	salt := make([]byte, SaltSize)

	_, err := rand.Read(salt)

	if err != nil {
		return nil, err
	}

	return salt, nil
}
