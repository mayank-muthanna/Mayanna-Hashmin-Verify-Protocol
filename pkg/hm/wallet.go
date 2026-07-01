package hm

type WalletField struct {
	VarName string
	Opening ValueOpening
}

type WalletCredential struct {
	PublicCredential Credential
	Fields           []WalletField
}
