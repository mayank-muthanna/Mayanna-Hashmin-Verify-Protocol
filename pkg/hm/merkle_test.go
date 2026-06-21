package hm

import "testing"

func TestMerkleLeavesProduceSameRoot(t *testing.T) {

	leaves := [][]byte{
		HashLeaf([]byte("A")),
		HashLeaf([]byte("B")),
		HashLeaf([]byte("C")),
		HashLeaf([]byte("D")),
	}

	root1 := BuildMerkleRoot(leaves)
	root2 := BuildMerkleRoot(leaves)

	if string(root1) != string(root2) {
		t.Fatalf("Same leaves are providing different root hash when sent through the same function twice")
	}

}
