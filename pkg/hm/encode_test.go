package hm

import (
	"testing"
)

func TestEncodeBool(t *testing.T) {
	encoded, err := EncodeValue(Value{
		Kind: Bool,
		Raw:  true,
	})

	if err != nil {
		t.Fatal(err)
	}

	if encoded.Bits[0] != 1 {
		t.Fatalf("expected TRUE to be encoded as Bit '1' But received otherwise")
	}
}

func TestEncodeInt8(t *testing.T) {
	encoded, err := EncodeValue(Value{
		Kind: Int8,
		Raw:  int64(18),
	})

	if err != nil {
		t.Fatal(err)
	}

	expected := []uint8{0, 0, 0, 1, 0, 0, 1, 0}

	for i := range expected {
		if encoded.Bits[i] != expected[i] {
			t.Fatalf("wrong bit at index %d", i)
		}
	}

}

func TestEncodeInt8RejectsOutOfRange(t *testing.T) {
	_, err := EncodeValue(Value{
		Kind: Int8,
		Raw:  int64(128),
	})

	if err == nil {
		t.Fatal("Expected Error, but function suceeded for int8, where given value was 128")
	}
}
