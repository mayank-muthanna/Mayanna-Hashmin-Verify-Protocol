package mpc

func Prove(circuit *Circuit, witness *Witness) (Proof, error) {

	views, err := EvaluateCircuit(circuit, witness)
	if err != nil {
		return Proof{}, err
	}

	transcript := NewTranscript(views)

	challenge, err := RandomChallenge()
	if err != nil {
		return Proof{}, err
	}

	proof := Proof{
		Challenge: challenge,
	}

	if challenge == Open01 {
		proof.View1 = transcript.Views[0]
		proof.View2 = transcript.Views[1]
	}

	if challenge == Open12 {
		proof.View1 = transcript.Views[1]
		proof.View2 = transcript.Views[2]
	}

	if challenge == Open20 {
		proof.View1 = transcript.Views[2]
		proof.View2 = transcript.Views[1]
	}

	return proof, nil

}
