package mpc

type OpenedView struct {
	Player int
	View   View
}

type RoundProof struct {
	Commitments [3]Commitment
	Challenge   Challenge

	View1 OpenedView
	View2 OpenedView
}

type Proof struct {
	Rounds []RoundProof
}
