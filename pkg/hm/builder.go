package hm

type CredentialBuilder struct {
	Schema Schema
	Values map[string]Value
}

func NewCredentialBuilder(schema Schema) CredentialBuilder {

	return CredentialBuilder{
		Schema: schema,
		Values: make(map[string]Value),
	}

}
