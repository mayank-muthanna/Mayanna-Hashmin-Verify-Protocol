package hm

func BuildFieldRoot(commitment ValueCommitment) []byte {

	if commitment.Kind == String {
		return HashLeaf(commitment.StringCommitment)
	}

	var leaves [][]byte

	for _, bitCommitment := range commitment.BitCommitments {

		leaf := HashLeaf(bitCommitment.Commitment)

		leaves = append(leaves, leaf)
	}

	merkleRoot := BuildMerkleRoot(leaves)

	return merkleRoot
}
