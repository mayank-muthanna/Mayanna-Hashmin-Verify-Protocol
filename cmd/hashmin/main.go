package main

import (
	"fmt"

	"github.com/mayank-muthanna/Mayanna-Hashmin-Verify-Protocol/pkg/hm"
)

func main() {
	value := hm.Value{
		Kind: hm.Int8,
		Raw:  int64(18),
	}

	encoded, err := hm.EncodeValue(value)

	if err != nil {
		panic(err)
	}

	fmt.Println("Mayanna HashMin Verify Protocol")
	fmt.Println("Kind: ", encoded.Kind)
	fmt.Println("Bits: ", encoded.Bits)
}
