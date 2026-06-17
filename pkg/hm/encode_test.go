package hm

import "testing"

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
