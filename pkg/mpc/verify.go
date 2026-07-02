package mpc

func Verify(circuit *Circuit, proof Proof) bool {

	for _, round := range proof.Rounds {
		if !VerifyRound(circuit, round) {
			return false
		}
	}

	return true
}

func VerifyRound(circuit *Circuit, round RoundProof) bool {

	challenge := FiatShamir(round.Commitments)

	if challenge != round.Challenge {
		return false
	}

	if !VerifyViews(circuit, round) {
		return false
	}

	return true
}

func OpenPlayer(_challenge Challenge) (uint8, uint8) {

	if _challenge == Open01 {
		return 0, 1
	}

	if _challenge == Open12 {
		return 1, 2
	}

	return 2, 0

}

func VerifyViews(circuit *Circuit, round RoundProof) bool {

	player1, player2 := OpenPlayer(
		round.Challenge,
	)

	if round.View1.Player != int(player1) {
		return false
	}

	if round.View2.Player != int(player2) {
		return false
	}

	if !verifySingleView(
		circuit,
		round.View1.View,
	) {
		return false
	}

	if !verifySingleView(
		circuit,
		round.View2.View,
	) {
		return false
	}

	return true
}

func verifySingleView(circuit *Circuit, view View) bool {

	for _, gate := range circuit.Gates {

		switch gate.GateType {

		case XOR:

			expected := view.Wires[gate.InputA] ^ view.Wires[gate.InputB]
			actual := view.Wires[gate.Output]

			if expected != actual {
				return false
			}

		case NOT:

			expected := view.Wires[gate.InputA] ^ 1

			actual := view.Wires[gate.Output]

			if expected != actual {
				return false
			}

		case AND:

			if _, exists := view.Wires[gate.Output]; !exists {
				return false
			}

		default:
			return false
		}
	}

	return true
}
