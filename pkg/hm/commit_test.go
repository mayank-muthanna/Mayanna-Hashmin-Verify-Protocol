package hm

import "testing"

func TestCommitBitVerifies(t *testing.T) {
	salt, err := RandomSalt()

	if err != nil {
		t.Fatal(err)
	}

	commonTestingBit := uint8(1)

	commitment, err := CommitBit(commonTestingBit, salt)

	if err != nil {
		t.Fatal(err)
	}

	opening := BitOpening{
		0,
		commonTestingBit,
		salt,
	}

	verified := VerifyBitOpening(commitment, opening)

	if !verified {
		t.Fatalf("Commitments not verified even with common values")
	}

}
