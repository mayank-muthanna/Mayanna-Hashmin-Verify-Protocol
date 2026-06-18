package hm

import (
	"crypto/sha256"
)

func Hash(domain string, parts ...[]byte) []byte {
	h := sha256.New()

	h.Write([]byte(domain))

	for _, part := range parts {
		h.Write(part)
	}

	return h.Sum(nil)
}
