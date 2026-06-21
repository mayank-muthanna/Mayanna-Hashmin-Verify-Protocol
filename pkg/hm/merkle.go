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
