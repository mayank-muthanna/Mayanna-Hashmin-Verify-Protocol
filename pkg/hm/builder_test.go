package hm

import "testing"

func TestCreateCredentialBuilder(t *testing.T) {

	schema := Schema{
		Name: "AADHAAR",
	}

	builder := NewCredentialBuilder(schema)

	if builder.Schema.Name != "AADHAAR" {
		t.Fatal()
	}

	if len(builder.Values) != 0 {
		t.Fatal()
	}

}
