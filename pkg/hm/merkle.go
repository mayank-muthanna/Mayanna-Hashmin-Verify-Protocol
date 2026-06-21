package hm

func HashNode(left []byte, right []byte) []byte {
	return Hash(
		"hm:merkle:node:v1",
		left,
		right,
	)
}

func HashLeaf(data []byte) []byte {
	return Hash(
		"hm:merkle:leaf:v1",
		data,
	)
}

func BuildNextLevel(currentLevel [][]byte) [][]byte {

	var nextLevel [][]byte

	for i := 0; i < len(currentLevel); i = +2 {

		left := currentLevel[i]

		var right []byte

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

func BuildMerkleRoot(leaves [][]byte) []byte {

	if len(leaves) == 0 {
		return nil
	}

	currentLevel := leaves

	for len(currentLevel) > 1 { //Making sure that there are alteast minimum 2 leaves
		currentLevel = BuildNextLevel(currentLevel)
	}

	return currentLevel[0]

}
