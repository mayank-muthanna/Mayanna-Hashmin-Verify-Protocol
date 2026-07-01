package hm

type Schema struct {
	SchemaName        string
	IdentifyingSecret string
	Fields            []SchemaField
	fieldMap          map[string]SchemaField // for easier searching of field values O(1) instead of looping
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
