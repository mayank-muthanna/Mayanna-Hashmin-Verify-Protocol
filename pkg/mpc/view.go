package mpc

type View struct {
	Player     int
	Wires      map[WireID]uint8
	Randomness []byte
}

func NewView(player int) *View {

	return &View{
		Player:     player,
		Wires:      make(map[WireID]uint8),
		Randomness: make([]byte, 0),
	}

}
