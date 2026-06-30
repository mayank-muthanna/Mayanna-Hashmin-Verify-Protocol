package hm

import "testing"

func TestMerkleLeavesProduceSameRoot(t *testing.T) {

	leaves := []Root{
		HashLeaf(Root("A")),
		HashLeaf(Root("B")),
		HashLeaf(Root("C")),
		HashLeaf(Root("D")),
	}

	root1 := BuildMerkleRoot(leaves)
	root2 := BuildMerkleRoot(leaves)

	if string(root1) != string(root2) {
		t.Fatalf("Same leaves are providing different root hash when sent through the same function twice")
	}

}

func TestMerkleProofVerifies(t *testing.T) {

	leaves := []Root{
		HashLeaf([]byte("A")),
		HashLeaf([]byte("B")),
		HashLeaf([]byte("C")),
		HashLeaf([]byte("D")),
	}

	root := BuildMerkleRoot(leaves)
	proof := BuildMerkleProof(leaves, 1)

	verifyRoots := VerifyMerkleProof(proof, root)

	if !verifyRoots {
		t.Fatal("Proof should verify, it did not")
	}

}

func TestMerkleProofRejectsWrongRoot(t *testing.T) {

	leaves := []Root{
		HashLeaf([]byte("A")),
		HashLeaf([]byte("B")),
		HashLeaf([]byte("C")),
		HashLeaf([]byte("D")),
	}

	root := BuildMerkleRoot(leaves)

	proof := BuildMerkleProof(leaves, 1)

	root[0] ^= 1 // Corrupt the root

	if VerifyMerkleProof(proof, root) {
		t.Fatal("proof should fail")
	}
}

func TestMerkleProofRejectsTamperedSibling(t *testing.T) {

	leaves := []Root{
		HashLeaf([]byte("A")),
		HashLeaf([]byte("B")),
		HashLeaf([]byte("C")),
		HashLeaf([]byte("D")),
	}

	root := BuildMerkleRoot(leaves)

	proof := BuildMerkleProof(leaves, 1)

	proof.Steps[0].Sibling[0] ^= 1

	if VerifyMerkleProof(proof, root) {
		t.Fatal("tampered proof should fail")
	}
}

func TestEveryLeafProofVerifies(t *testing.T) {

	leaves := []Root{
		HashLeaf([]byte("A")),
		HashLeaf([]byte("B")),
		HashLeaf([]byte("C")),
		HashLeaf([]byte("D")),
		HashLeaf([]byte("E")),
	}

	root := BuildMerkleRoot(leaves)

	for i := range leaves {

		proof := BuildMerkleProof(leaves, i)

		if !VerifyMerkleProof(proof, root) {
			t.Fatalf("proof failed for leaf %d", i)
		}
	}
}
