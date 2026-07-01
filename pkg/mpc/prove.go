package mpc

func ProveRound(circuit *Circuit, witness *Witness) (RoundProof, error) {

	views, err := EvaluateCircuit(circuit, witness)
	if err != nil {
		return RoundProof{}, err
	}

	transcript := NewTranscript(views)

	commitments, err := CommitViews(views)

	if err != nil {
		return RoundProof{}, err
	}

	challenge := FiatShamir(commitments)

	proof := RoundProof{
		Commitments: commitments,
		Challenge:   challenge,
	}

	if challenge == Open01 {
		proof.View1 = OpenedView{Player: 0, View: transcript.Views[0]}
		proof.View2 = OpenedView{Player: 1, View: transcript.Views[1]}
	}

	if challenge == Open12 {
		proof.View1 = OpenedView{Player: 1, View: transcript.Views[1]}
		proof.View2 = OpenedView{Player: 2, View: transcript.Views[2]}
	}

	if challenge == Open20 {
		proof.View1 = OpenedView{Player: 2, View: transcript.Views[2]}
		proof.View2 = OpenedView{Player: 0, View: transcript.Views[0]}
	}

	return proof, nil

}

const Repetitions = 80

func Prove(circuit *Circuit, witness *Witness) (Proof, error) {

	proof := Proof{}

	for i := 0; i < Repetitions; i++ {

		round, err := ProveRound(
			circuit,
			witness,
		)

		if err != nil {
			return Proof{}, err
		}

		proof.Rounds = append(
			proof.Rounds,
			round,
		)
	}

	return proof, nil
}
