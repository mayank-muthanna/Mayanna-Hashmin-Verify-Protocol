package mpc

type RoundProof struct {
	Challenge Challenge

	View1 View
	View2 View
}

type Proof struct {
	Rounds []RoundProof
}
