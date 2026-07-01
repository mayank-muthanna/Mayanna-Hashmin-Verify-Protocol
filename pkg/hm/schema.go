package hm

type SchemaField struct {
	VarName string
	Kind    Kind
}

type Schema struct {
	SchemaName        string
	IdentifyingSecret string
	Fields            []SchemaField
	fieldMap          map[string]SchemaField // for easier searching of field values O(1) instead of looping
}
