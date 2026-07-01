package hm

type SchemaField struct {
	Name string
	Kind Kind
}

type Schema struct {
	Name              string
	IdentifyingSecret string
	Fields            []SchemaField
	fieldMap          map[string]SchemaField // for easier searching of field values O(1) instead of looping
}
