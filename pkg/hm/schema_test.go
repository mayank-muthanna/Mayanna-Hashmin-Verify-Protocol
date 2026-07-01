package hm

import "testing"

func TestCreateSchema(t *testing.T) {

	schema := Schema{
		SchemaName:        "AADHAAR_AUTHORITY",
		IdentifyingSecret: "aadhaar_number",
		Fields: []SchemaField{
			{
				VarName: "name",
				Kind:    String,
			},
			{
				VarName: "age",
				Kind:    Int8,
			},
			{
				VarName: "student",
				Kind:    Bool,
			},
		},
	}

	if schema.SchemaName != "AADHAAR_AUTHORITY" {
		t.Fatal("incorrect schema name")
	}

	if schema.IdentifyingSecret != "aadhaar_number" {
		t.Fatal("incorrect identifying secret")
	}

	if len(schema.Fields) != 3 {
		t.Fatal("expected 3 fields")
	}

}
