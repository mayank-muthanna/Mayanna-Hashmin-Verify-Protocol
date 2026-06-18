package hm

import (
	"crypto/rand"
	"fmt"
)

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

func CommitBit(bit uint8, salt []byte) ([]byte, error) {

	if bit != 0 && bit != 1 {
		return nil, fmt.Errorf("%d should be {0, 1} ", bit)
	}

	if len(salt) != SaltSize {
		return nil, fmt.Errorf("SaltSize cannot exceed %d bits, received %s", SaltSize, string(salt))
	}

	domainName := string("hm:bit:v1")

	return Hash(
		domainName,
		[]byte{bit},
		salt,
	), nil
}

func VerifyBitOpening(commitment []byte, opening BitOpening) bool {

	commit_derived_from_opening, err := CommitBit(opening.Bit, opening.Salt)

	if err != nil {
		return false
	}

	if string(commit_derived_from_opening) != string(commitment) {
		return false
	}

	return true
}
