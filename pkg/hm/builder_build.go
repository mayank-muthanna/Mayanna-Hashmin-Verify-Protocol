package hm

import "fmt"

func (builder *CredentialBuilder) Build() (WalletCredential, error) {

	credential := Credential{
		SchemaName: builder.Schema.SchemaName,
	}

	wallet := WalletCredential{
		PublicCredential: credential,
	}

	for _, field := range builder.Schema.Fields {

		credentialValue, exists := builder.Values[field.VarName]

		if !exists {
			return WalletCredential{}, fmt.Errorf("Missing value of %s", field.VarName)
		}

		encodedValue, err := EncodeValue(credentialValue.Value)
		if err != nil {
			return WalletCredential{}, err
		}

		commitment, opening, err := CommitEncodedValue(encodedValue)
		if err != nil {
			return WalletCredential{}, err
		}

		field := CredentialField{
			VarName:    credentialValue.Field.VarName,
			Commitment: commitment,
		}

		wallet.PublicCredential.Fields = append(
			wallet.PublicCredential.Fields,
			field,
		)

		wallet.Fields = append(
			wallet.Fields,
			WalletField{
				VarName: credentialValue.Field.VarName,
				Opening: opening,
			},
		)
	}

	wallet.PublicCredential.Root = BuildCredentialRoot(
		wallet.PublicCredential.Fields,
	)

	return wallet, nil
}
