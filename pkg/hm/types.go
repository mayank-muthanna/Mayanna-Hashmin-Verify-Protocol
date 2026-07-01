package hm

type Kind string

const (
	Bool   Kind = "bool"
	Byte   Kind = "byte"
	String Kind = "string"

	Int4  Kind = "int4"
	Int8  Kind = "int8"
	Int16 Kind = "int16"
	Int32 Kind = "int32"
	Int64 Kind = "int64"
)

type Value struct {
	Kind Kind
	Raw  any
}

type EncodedValue struct {
	Kind  Kind
	Bytes []byte
	Bits  []uint8
}

type SchemaField struct {
	VarName string
	Kind    Kind
}

type CredentialValue struct {
	Field SchemaField
	Value Value
}

//Readablity Friendly types to be used in fxns

type Root []byte

type ProofStep struct {
	Sibling Root
	IsLeft  bool
}

type MerkleProof struct {
	LeafIndex int
	Leaf      Root
	Steps     []ProofStep
}
