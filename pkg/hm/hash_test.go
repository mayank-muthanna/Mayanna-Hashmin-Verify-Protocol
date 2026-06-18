package hm

import "testing"

func TestHashSimilarity(t *testing.T) {
	similarityChkStr := string("hello")
	a := Hash("test", []byte(similarityChkStr))
	b := Hash("test", []byte(similarityChkStr))

	if string(a) != string(b) {
		t.Fatalf("Hash function f(x) with same inputs is resulting in different output")
	}
}
