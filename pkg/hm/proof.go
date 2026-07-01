package hm

type Proof struct {
	CircuitProof    CircuitProof
	CommitmentProof CommitmentProof
}

type CircuitProof struct {
}

type CommitmentProof struct {
	Fields []CommitmentFieldProof
}

type CommitmentFieldProof struct {
	VarName string
	Opening ValueOpening
}
