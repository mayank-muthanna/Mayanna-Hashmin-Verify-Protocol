package hm

import "testing"

func TestCreateSchema(t *testing.T) {

	schema := Schema{
		Name:              "AADHAAR_AUTHORITY",
		IdentifyingSecret: "aadhaar_number",
		Fields: []SchemaField{
			{
				Name: "name",
				Kind: String,
			},
			{
				Name: "age",
				Kind: Int8,
			},
			{
				Name: "student",
				Kind: Bool,
			},
		},
	}

	if schema.Name != "AADHAAR_AUTHORITY" {
		t.Fatal("incorrect schema name")
	}

	if schema.IdentifyingSecret != "aadhaar_number" {
		t.Fatal("incorrect identifying secret")
	}

	if len(schema.Fields) != 3 {
		t.Fatal("expected 3 fields")
	}

}
