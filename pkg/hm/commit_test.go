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

func TestWrongBitCommitVerifyFails(t *testing.T) {
	salt, err := RandomSalt()

	if err != nil {
		t.Fatal(err)
	}

	TestingBit := uint8(1)
	TestingBitInverted := uint8(0)

	commitment, err := CommitBit(TestingBit, salt)

	if err != nil {
		t.Fatal(err)
	}

	opening := BitOpening{
		0,
		TestingBitInverted,
		salt,
	}

	verified := VerifyBitOpening(commitment, opening)

	if verified {
		t.Fatalf("Commitments are being verified even with differing values")
	}
}

func TestWrongSaltVerificationFails(t *testing.T) {

	salt, err := RandomSalt()

	if err != nil {
		t.Fatal(err)
	}

	wrongSalt, err := RandomSalt()

	if err != nil {
		t.Fatal(err)
	}

	CommonTestingBit := uint8(1)

	commitment, err := CommitBit(CommonTestingBit, wrongSalt)

	if err != nil {
		t.Fatal(err)
	}

	opening := BitOpening{
		0,
		CommonTestingBit,
		salt,
	}

	if VerifyBitOpening(commitment, opening) {
		t.Fatalf("Commitments are being verified even with differing salts")
	}

}

func TestCommitEncodedInt8Value(t *testing.T) {

	encoded, err := EncodeValue(Value{
		Kind: Int8,
		Raw:  int64(18),
	})

	if err != nil {
		t.Fatal(err)
	}

	commitment, opening, err := CommitEncodedValue(encoded)

	if err != nil {
		t.Fatal(err)
	}

	if commitment.Kind != Int8 {
		t.Fatal("Expected Int8 commitment")
	}

	if len(commitment.BitCommitments) != 8 {
		t.Fatal("Expected 8 bit commitments")
	}

	if len(opening.BitOpenings) != 8 {
		t.Fatal("Expected 8 bit openings")
	}

}
