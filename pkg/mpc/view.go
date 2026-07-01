package mpc

type View struct {
	Player       int
	InputShares  map[WireID]uint8
	OutputShares map[WireID]uint8
	Randomness   []byte
}

func NewView(player int) *View {

	return &View{
		Player:       player,
		InputShares:  make(map[WireID]uint8),
		OutputShares: make(map[WireID]uint8),
	}

}
