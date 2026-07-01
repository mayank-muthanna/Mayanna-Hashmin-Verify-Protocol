package hm

type CompiledPolicy struct {
	CircuitPredicates    []Predicate
	CommitmentPredicates []Predicate
}

func CompilePolicy(policy Policy) CompiledPolicy {

	compiled := CompiledPolicy{}

	for _, predicate := range policy.Predicates {

		if predicate.Field.Kind == String {

			compiled.CommitmentPredicates = append(
				compiled.CommitmentPredicates,
				predicate,
			)

			continue
		}

		compiled.CircuitPredicates = append(
			compiled.CircuitPredicates,
			predicate,
		)
	}

	return compiled
}
