package hm

import (
	"fmt"
)

type CredentialBuilder struct {
	Schema Schema
	Values map[string]CredentialValue
}

func NewCredentialBuilder(schema Schema) *CredentialBuilder {

	return &CredentialBuilder{
		Schema: schema,
		Values: make(map[string]CredentialValue),
	}

}

func (builder *CredentialBuilder) Add(varName string, raw any) error {

	field, err := builder.findField(varName)

	if err != nil {
		return err
	}

	//validate before storing
	if err := validateValue(field, raw); err != nil {
		return err
	}

	builder.storeValue(field, raw)

	return nil
}

// priv fxns to be called by Add
func (builder *CredentialBuilder) findField(varName string) (SchemaField, error) {

	field, exists := builder.Schema.fieldMap[varName]

	if !exists {
		return SchemaField{}, fmt.Errorf("field %s does not exists", varName)
	}

	return field, nil

}

func (builder *CredentialBuilder) storeValue(field SchemaField, raw any) {

	builder.Values[field.VarName] = CredentialValue{
		Field: field,
		Value: Value{
			Kind: field.Kind,
			Raw:  raw,
		},
	}

}

func validateValue(field SchemaField, raw any) error {

	switch field.Kind {

	case Bool:
		_, ok := raw.(bool)
		if !ok {
			return fmt.Errorf("%s expects bool", field.VarName)
		}

	case String:
		_, ok := raw.(string)
		if !ok {
			return fmt.Errorf("%s expects string", field.VarName)
		}

	case Byte:
		_, ok := raw.(byte)
		if !ok {
			return fmt.Errorf("%s expects byte", field.VarName)
		}

	case Int4, Int8, Int16, Int32, Int64:
		_, ok := raw.(int64)
		if !ok {
			return fmt.Errorf("%s expects int64", field.VarName)
		}

	default:
		return fmt.Errorf("unsupported type %s", field.Kind)
	}

	return nil
}
