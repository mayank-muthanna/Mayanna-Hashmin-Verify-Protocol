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

// Test if Addition of Variables
func TestBuilderAdd(t *testing.T) {

	schema := NewSchema(
		"AADHAAR",
		"aadhaar_number",
		[]SchemaField{
			{VarName: "age", Kind: Int8},
		},
	)

	builder := NewCredentialBuilder(schema)

	err := builder.Add("age", int64(18))

	if err != nil {
		t.Fatal(err)
	}

	if len(builder.Values) != 1 {
		t.Fatal("expected one value")
	}
}

// Addition of non-schema-declared variables should fail
func TestBuilderRejectUnknownField(t *testing.T) {

	schema := NewSchema(
		"AADHAAR",
		"aadhaar_number",
		nil,
	)

	builder := NewCredentialBuilder(schema)

	if builder.Add("hello", int64(1)) == nil {
		t.Fatal("expected error")
	}
}

// Wrong Field type should fail
func TestBuilderRejectWrongType(t *testing.T) {

	schema := NewSchema(
		"AADHAAR",
		"aadhaar_number",
		[]SchemaField{
			{VarName: "age", Kind: Int8},
		},
	)

	builder := NewCredentialBuilder(schema)

	if builder.Add("age", "eighteen") == nil {
		t.Fatal("expected type error")
	}
}
