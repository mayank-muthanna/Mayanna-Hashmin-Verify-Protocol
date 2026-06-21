package hm

import "testing"

func TestBuildInt8FieldRoot(t *testing.T) {

	encoded, err := EncodeValue(Value{
		Kind: Int8,
		Raw:  int64(18),
	})

	if err != nil {
		t.Fatal(err)
	}

	commitment, _, err := CommitEncodedValue(encoded)

	if err != nil {
		t.Fatal(err)
	}

	root1 := BuildFieldRoot(commitment)
	root2 := BuildFieldRoot(commitment)

	if string(root1) != string(root2) {
		t.Fatalf("Same commitments should habe given same roots but received otherwise")
	}

}

func TestBuildStringFieldRoot(t *testing.T) {

	encoded, err := EncodeValue(Value{
		Kind: String,
		Raw:  string("MayankMuthanna"),
	})

	if err != nil {
		t.Fatal(err)
	}

	commitment, _, err := CommitEncodedValue(encoded)

	if err != nil {
		t.Fatal(err)
	}

	root := BuildFieldRoot(commitment)

	if len(string(root)) == 0 {
		t.Fatal("Empty merkle root for string")
	}
}

func TestBuildCredentialRoot(t *testing.T) {

	ageEncoded, _ := EncodeValue(Value{
		Kind: Int8,
		Raw:  int64(18),
	})

	isStudentEncoded, _ := EncodeValue(Value{
		Kind: Bool,
		Raw:  true,
	})

	ageCommitment, _, _ := CommitEncodedValue(ageEncoded)
	isStudentBoolCommitment, _, _ := CommitEncodedValue(isStudentEncoded)

	fields := []CredentialFiled{
		{
			Name:       "age",
			Commitment: ageCommitment,
		},
		{
			Name:       "isStudent",
			Commitment: isStudentBoolCommitment,
		},
	}

	root1 := BuildCredentialRoot(fields)
	root2 := BuildCredentialRoot(fields)

	if string(root1) != string(root2) {
		t.Fatal("Same fields are producing different roots when sent thorugh function twice")
	}
}
