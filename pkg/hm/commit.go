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

type ValueCommitment struct {
	Kind Kind

	// Used for bool, byte, int4, int8, int16, int32, int64
	BitCommitments []BitCommitment

	// Used only for string
	StringCommitment []byte
}

type ValueOpening struct {
	Kind Kind

	// Used for bool, byte, int4, int8, int16, int32, int64
	BitOpenings []BitOpening

	// Used only for string
	StringSalt []byte
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

func CommitEncodedValue(encoded EncodedValue) (ValueCommitment, ValueOpening, error) {
	if encoded.Kind == String {
		return commitStringValue(encoded)
	}

	return commitBitsValue(encoded)
}

func commitBitsValue(encoded EncodedValue) (ValueCommitment, ValueOpening, error) {

	var bitCommitments_ []BitCommitment
	var bitOpenings_ []BitOpening

	for index, bit := range encoded.Bits {

		salt, err := RandomSalt()

		if err != nil {
			return ValueCommitment{}, ValueOpening{}, err
		}

		commitment, err := CommitBit(bit, salt)

		if err != nil {
			return ValueCommitment{}, ValueOpening{}, err
		}

		_bitCommitment := BitCommitment{
			Index:      index,
			Commitment: commitment,
		}

		_bitOpening := BitOpening{
			Index: index,
			Bit:   bit,
			Salt:  salt,
		}

		bitCommitments_ = append(bitCommitments_, _bitCommitment)
		bitOpenings_ = append(bitOpenings_, _bitOpening)
	}

	_valueCommitment := ValueCommitment{
		Kind:           encoded.Kind,
		BitCommitments: bitCommitments_,
	}

	_valueOpening := ValueOpening{
		Kind:        encoded.Kind,
		BitOpenings: bitOpenings_,
	}

	return _valueCommitment, _valueOpening, nil

}

func commitStringValue(encoded EncodedValue) (ValueCommitment, ValueOpening, error) {

	salt, err := RandomSalt()

	if err != nil {
		return ValueCommitment{}, ValueOpening{}, err
	}

	commitment_ := Hash(
		"hm:verify:v1",
		encoded.Bytes,
		salt,
	)

	valueCommitment_ := ValueCommitment{
		Kind:             String,
		StringCommitment: commitment_,
	}

	valueOpening_ := ValueOpening{
		Kind:       String,
		StringSalt: salt,
	}

	return valueCommitment_, valueOpening_, nil

}
