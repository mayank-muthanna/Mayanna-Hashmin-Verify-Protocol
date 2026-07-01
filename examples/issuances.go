package main

import (
	"fmt"

	"github.com/mayank-muthanna/Mayanna-Hashmin-Verify-Protocol/pkg/hm"
)

func main() {

	//1. CREATE SCHEMA

	schema := hm.NewSchema(
		"AADHAAR",
		"adhaar_authorirty",
		[]hm.SchemaField{
			{
				VarName: "name",
				Kind:    hm.String,
			},
			{
				VarName: "age",
				Kind:    hm.Int8,
			},
			{
				VarName: "student",
				Kind:    hm.Bool,
			},
		},
	)

	// 2. CREATE BUILDER

	builder := hm.NewCredentialBuilder(schema)

	// 3. ADD VALUES

	builder.Add("name", "MayankMuthanna")
	builder.Add("age", int64(18))
	builder.Add("student", true)

	// 4. CONSTRUCT WALLET

	wallet, err := builder.Build()

	if err != nil {
		panic(err)
	}

	// 5. PRINT RESULTS

	fmt.Println("Credential Root:")
	fmt.Printf("%x\n\n", wallet.PublicCredential.Root)

	fmt.Println("Fields:")

	for _, field := range wallet.PublicCredential.Fields {

		fmt.Println(field.VarName)
	}

	fmt.Println()

	fmt.Println("Stored Openings:")

	fmt.Println(len(wallet.Fields))

}
