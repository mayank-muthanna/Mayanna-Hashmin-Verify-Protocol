package hm

type SchemaField struct {
	Name string
	Kind Kind
}

type Schema struct {
	Name              string
	IdentifyingSecret string
	Fields            []SchemaField
}
