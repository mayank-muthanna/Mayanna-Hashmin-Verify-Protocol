package hm

import "bytes"

func HashNode(left Root, right Root) Root {
	return Hash(
		"hm:merkle:node:v1",
		left,
		right,
	)
}

func HashLeaf(data []byte) Root {
	return Hash(
		"hm:merkle:leaf:v1",
		data,
	)
}

// eg. 8 roots become 4 and 4 roots become 2
func BuildNextLevel(currentLevel []Root) []Root {

	var nextLevel []Root

	for i := 0; i < len(currentLevel); i += 2 {

		left := currentLevel[i]

		var right Root

		if i+1 < len(currentLevel) {
			right = currentLevel[i+1]
		} else {
			right = left
		}

		parent := HashNode(left, right)

		nextLevel = append(nextLevel, parent)

	}

	return nextLevel

}

// uses above function to build Merkle Tree
func BuildMerkleRoot(leaves []Root) Root {

	if len(leaves) == 0 {
		return nil
	}

	currentLevel := leaves

	//Making sure that we are sending atleast 2 roots to BuildNextLevel function so that it returns 1 and then stop
	for len(currentLevel) > 1 {
		currentLevel = BuildNextLevel(currentLevel)
	}

	return currentLevel[0]

}

// 0,1 -> 0 		2,3 -> 1 		4,5 -> 2
func ParentIndex(index int) int {
	return index / 2
}

// Finds index of sibling (if left in case 0,1: 1 -> gives 0 (left sibling) )
func SiblingIndex(index int) int {

	if index%2 == 0 {
		return index + 1
	}

	return index - 1

}

func FindSibling(currentLevel []Root, index int) (Root, bool) {

	siblingIndex := SiblingIndex(index)

	//If there are no available sibling in 2 pairs, return same root - for cases where merkle has odd number of leaves
	if siblingIndex >= len(currentLevel) {
		return currentLevel[index], false
	}

	return currentLevel[siblingIndex], true

}

func BuildMerkleProof(leaves []Root, index int) MerkleProof {

	proof := MerkleProof{
		LeafIndex: index,
		Leaf:      leaves[index],
	}

	currentLevel := leaves
	currentIndex := index

	for len(currentLevel) > 1 {

		sibling, exists := FindSibling(currentLevel, currentIndex)

		step := ProofStep{
			Sibling: sibling,
			IsLeft:  currentIndex%2 == 1,
		}

		//ingored for timebeing
		_ = exists

		proof.Steps = append(proof.Steps, step)

		currentIndex = ParentIndex(currentIndex)
		currentLevel = BuildNextLevel(currentLevel)

	}

	return proof
}

func VerifyMerkleProof(proof MerkleProof, expectedRoot Root) bool {

	current := proof.Leaf

	for _, step := range proof.Steps {

		if step.IsLeft {
			current = HashNode(step.Sibling, current)
		} else {
			current = HashNode(current, step.Sibling)
		}

	}

	var isConstructedMatchExpected bool

	if bytes.Equal(current, expectedRoot) {
		isConstructedMatchExpected = true
	} else {
		isConstructedMatchExpected = false
	}

	return isConstructedMatchExpected

}
