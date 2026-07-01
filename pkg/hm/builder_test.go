package hm

import "testing"

func TestCreateCredentialBuilder(t *testing.T) {

	schema := NewSchema(
		"ADHAAR_AUTHORITY",
		"adhaar_authority",
		[]SchemaField{
			{VarName: "name", Kind: String},
			{VarName: "age", Kind: Int8},
			{VarName: "student", Kind: Bool},
		},
	)

	field, ok := schema.fieldMap["age"]

	if !ok {
		t.Fatal("expected age field")
	}

	if field.Kind != Int8 {
		t.Fatal("incorrect field kind")
	}

	builder := NewCredentialBuilder(schema)

	if builder.Schema.SchemaName != "ADHAAR_AUTHORITY" {
		t.Fatal()
	}

	if len(builder.Values) != 0 {
		t.Fatal()
	}

}
