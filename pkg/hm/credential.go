package hm

type CredentialFiled struct {
	Name       string
	Commitment ValueCommitment
}

type Credential struct {
	Fields []CredentialFiled
	Root   Root
}

func BuildFieldRoot(commitment ValueCommitment) Root {

	if commitment.Kind == String {
		return HashLeaf(commitment.StringCommitment)
	}

	var leaves []Root

	for _, bitCommitment := range commitment.BitCommitments {

		leaf := HashLeaf(bitCommitment.Commitment)

		leaves = append(leaves, leaf)
	}

	merkleRoot := BuildMerkleRoot(leaves)

	return merkleRoot
}

func BuildCredentialRoot(fields []CredentialFiled) Root {

	var fieldRoots []Root

	for _, field := range fields {

		root := BuildFieldRoot(field.Commitment)
		fieldRoots = append(fieldRoots, root)
	}

	merkleRoot := BuildMerkleRoot(fieldRoots)

	return merkleRoot

}
