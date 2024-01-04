package tracker

import (
	"github.com/quasilyte/drum-rpg/edrum"
	"github.com/quasilyte/xm"
	"github.com/quasilyte/xm/xmfile"
)

type MixedTrack struct {
	Track         *Track
	Module        *xmfile.Module
	Stream        *xm.Stream
	InputChannels [4][]InputNote
	Duration      float64
}

type InputNote struct {
	Time       float64
	Instrument edrum.InstrumentKind
}

type MixerConfig struct {
	Track *Track
	Tempo int
	BPM   int
}

func Mix(config MixerConfig) (*MixedTrack, error) {
	t := config.Track

	tempo := config.Tempo
	if tempo == 0 {
		tempo = config.Track.Module.DefaultTempo
	}
	if tempo == 0 {
		tempo = 6
	}

	bpm := config.BPM
	if bpm == 0 {
		bpm = config.Track.Module.DefaultBPM
	}
	if bpm == 0 {
		bpm = 120
	}

	module := &xmfile.Module{}

	// Start with a shallow copy.
	// We'll re-allocate the necessary slices later.
	*module = *t.Module

	mt := &MixedTrack{
		Track:  t,
		Module: module,
	}

	// +1 is needed for an empty pattern.
	// That empty pattern will be used as a first pattern to
	// give the player some delay before the main part starts.
	module.NumPatterns = len(module.Patterns) + 1

	numEmptyRows := 64

	mixedPatterns := make([]xmfile.Pattern, module.NumPatterns)
	for i := range module.Patterns {
		p := &module.Patterns[i]
		mixed := &mixedPatterns[i]
		mixed.Rows = make([]xmfile.PatternRow, len(p.Rows))
		for j := range p.Rows {
			row := &p.Rows[j]
			mixedRow := &mixed.Rows[j]
			mixedRow.Notes = make([]uint16, len(row.Notes))
			for channelID, noteID := range row.Notes {
				n := module.Notes[noteID]
				kind := t.GetInstrumentKind(int(n.Instrument))
				if kind != edrum.UndefinedInstrument {
					continue
				}
				mixedRow.Notes[channelID] = noteID
			}
		}
	}
	// Fill the empty pattern.
	{
		numRows := numEmptyRows
		p := xmfile.Pattern{}
		p.Rows = make([]xmfile.PatternRow, numRows)
		for i := range p.Rows {
			p.Rows[i].Notes = make([]uint16, module.NumChannels)
		}
		mixedPatterns[len(mixedPatterns)-1] = p
	}
	// Make empty pattern go first.
	{
		patternOrder := make([]uint8, len(module.PatternOrder)+1)
		copy(patternOrder[1:], module.PatternOrder)
		patternOrder[0] = uint8(len(mixedPatterns) - 1)
		module.PatternOrder = patternOrder
	}
	module.Patterns = mixedPatterns

	// Build the input channels by traversing the patterns in their play order.
	secondsPerRow := calcSecondsPerRow(tempo, float64(bpm))
	rowIndex := numEmptyRows
	for _, patternIndex := range t.Module.PatternOrder {
		p := &t.Module.Patterns[patternIndex]
		for j := range p.Rows {
			row := &p.Rows[j]
			rowTime := calcRowTime(rowIndex, secondsPerRow)
			for _, noteID := range row.Notes {
				n := module.Notes[noteID]
				kind := t.GetInstrumentKind(int(n.Instrument))
				if kind == edrum.UndefinedInstrument {
					continue
				}
				channelID := kind.Channel()
				mt.InputChannels[channelID] = append(mt.InputChannels[channelID], InputNote{
					Instrument: kind,
					Time:       rowTime,
				})
			}
			rowIndex++
		}
	}

	stream := xm.NewStream()
	err := stream.LoadModule(module, xm.LoadModuleConfig{
		LinearInterpolation: true,
		Tempo:               uint(tempo),
		BPM:                 uint(bpm),
	})
	if err != nil {
		return nil, err
	}
	mt.Stream = stream

	// This is an approximation.
	mt.Duration = calcRowTime(rowIndex, secondsPerRow)

	return mt, nil
}

func calcSecondsPerRow(ticksPerRow int, bpm float64) float64 {
	ticksPerSecond := bpm * 0.4
	return 1 / (ticksPerSecond / float64(ticksPerRow))
}

func calcRowTime(rowIndex int, secondsPerRow float64) float64 {
	return float64(rowIndex) * secondsPerRow
}
