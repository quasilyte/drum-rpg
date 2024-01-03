package edrum

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
)

type InstrumentKind int

const (
	UndefinedInstrument InstrumentKind = iota

	BassInstrument // aka Kick Drum

	SnareInstrument

	LeftTomInstrument  // aka High Tom
	RightTomInstrument // aka Mid Tom
	FloorTomInstrument // aka Low Tom

	ClosedHiHatInstrument
	OpenHiHatInstrument

	LeftCymbalInstrument  // aka Crash Cymbal
	RightCymbalInstrument // aka Ride Cymbal

	_numInstruments
)

type Instrument struct {
	SampleSelection SampleSelection

	sampleMap [128][]resource.AudioID
}

type Sample struct {
	MinVelocity float64
	MaxVelocity float64

	AudioID resource.AudioID

	Name string
}

type SampleSelection int

const (
	FirstSampleSelection SampleSelection = iota
	RandomSampleSelection
)

func NewInstrument(sampleSelection SampleSelection, samples []Sample) *Instrument {
	inst := &Instrument{
		SampleSelection: sampleSelection,
	}

	for _, s := range samples {
		minVelocity := uint8(gmath.Clamp(s.MinVelocity*127.0, 0, 127))
		maxVelocity := uint8(gmath.Clamp(s.MaxVelocity*127.0, 0, 127))
		if s.MaxVelocity == 0 {
			maxVelocity = 127
		}
		numSteps := int(maxVelocity) - int(minVelocity)
		if numSteps <= 0 {
			panic("instrument with invalid min/max velocity")
		}
		for v := int(minVelocity); v <= int(maxVelocity); v++ {
			inst.sampleMap[v] = append(inst.sampleMap[v], s.AudioID)
		}
	}

	return inst
}

func (inst *Instrument) GetSamples(velocity uint8) []resource.AudioID {
	return inst.sampleMap[velocity&0x7f]
}
