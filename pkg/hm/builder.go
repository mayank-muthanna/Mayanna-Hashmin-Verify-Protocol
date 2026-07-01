package hm

type CredentialBuilder struct {
	Schema Schema
	Values map[string]CredentialValue
}

func NewSchema(name string, identifyingSecret string, fields []SchemaField) Schema {

	_fieldMap := make(map[string]SchemaField)

	for _, field := range fields {
		_fieldMap[field.VarName] = field
	}

	return Schema{
		SchemaName:        name,
		IdentifyingSecret: identifyingSecret,
		Fields:            fields,
		fieldMap:          _fieldMap,
	}

}

func NewCredentialBuilder(schema Schema) CredentialBuilder {

	return CredentialBuilder{
		Schema: schema,
		Values: make(map[string]CredentialValue),
	}

}
