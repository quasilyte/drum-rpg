package edrum

type Kit struct {
	Name string

	NoteMap NoteMap

	InstrumentMap [_numInstruments]*Instrument
}

func NewKit(m NoteMap, instruments map[InstrumentKind]*Instrument) *Kit {
	kit := &Kit{NoteMap: m}

	for k, inst := range instruments {
		kit.InstrumentMap[k] = inst
	}

	return kit
}
