package hm

import "fmt"

type PredciateComparision string

const (
	Equal       PredciateComparision = "eq"
	GreaterThan PredciateComparision = "gt"
	LessThan    PredciateComparision = "lt"
)

type Predicate struct {
	Field         SchemaField
	PredicateType PredciateComparision
	Value         Value
}

type Policy struct {
	Schema     Schema
	Predicates []Predicate
}

func NewPolicy(_schema Schema) *Policy {

	return &Policy{
		Schema: _schema,
	}

}

func (policy *Policy) AddEqual(varName string, raw any) error {

	field, err := policy.findField(varName)

	if err != nil {
		return err
	}

	return policy.addPredicate(field, Equal, raw)

}

func (policy *Policy) AddGreaterThan(varName string, value int64) error {

	field, err := policy.findField(varName)

	if err != nil {
		return err
	}

	return policy.addPredicate(field, GreaterThan, value)

}

func (policy *Policy) AddLessThan(varName string, value int64) error {

	field, err := policy.findField(varName)

	if err != nil {
		return err
	}

	return policy.addPredicate(field, LessThan, value)

}

// finds if field exists in policy from its attached schema
func (policy *Policy) findField(varName string) (SchemaField, error) {

	field, exists := policy.Schema.fieldMap[varName]

	if !exists {
		return SchemaField{}, fmt.Errorf("%s field not found in schema", varName)
	}

	return field, nil

}

// adds predicates to policy
func (policy *Policy) addPredicate(field SchemaField, predicateType PredciateComparision, raw any) error {

	err := validateValue(field, raw)
	if err != nil {
		return err
	}

	_value := Value{
		Kind: field.Kind,
		Raw:  raw,
	}

	_predicateToAdd := Predicate{
		Field:         field,
		PredicateType: predicateType,
		Value:         _value,
	}

	policy.Predicates = append(policy.Predicates, _predicateToAdd)

	return nil

}
