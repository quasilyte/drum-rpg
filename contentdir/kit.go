package contentdir

type DrumKitInstrument struct {
	Key string
}

type DrumKit struct {
	Name string

	Instruments map[string]DrumKitInstrument
}
