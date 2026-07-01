package mpc

type Witness struct {
	Wires map[WireID]uint8
}

func NewWitness() *Witness {
	return &Witness{
		Wires: make(map[WireID]uint8),
	}

}

func (w *Witness) Set(wire WireID, value uint8) {

	if value != 0 {
		value = 1
	}

	w.Wires[wire] = value

}

func (w *Witness) Get(wire WireID) uint8 {
	return w.Wires[wire]
}
