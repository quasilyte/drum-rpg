package edrum

type NoteMap [128]InstrumentKind

var RolandTD02kvMap = NoteMap{
	22: ClosedHiHatInstrument,
	26: OpenHiHatInstrument,
	36: BassInstrument,
	38: SnareInstrument,
	42: ClosedHiHatInstrument,
	43: FloorTomInstrument,
	45: RightTomInstrument,
	46: OpenHiHatInstrument,
	48: LeftTomInstrument,
	49: LeftCymbalInstrument,
	51: RightCymbalInstrument,
	55: LeftCymbalInstrument,
}
