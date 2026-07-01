package mpc

func Verify(proof Proof) bool {

	for _, round := range proof.Rounds {

		if !VerifyRound(round) {
			return false
		}
	}

	return true
}

func VerifyRound(round RoundProof) bool {

	challenge := FiatShamir(round.Commitments)

	if challenge != round.Challenge {
		return false
	}

	return true
}
