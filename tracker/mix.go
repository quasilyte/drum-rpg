package tracker

import (
	"github.com/quasilyte/xm"
	"github.com/quasilyte/xm/xmfile"
)

type MixedTrack struct {
	Track  *Track
	Module *xmfile.Module
	Stream *xm.Stream
}

type MixerConfig struct {
	Track *Track
}

func Mix(config MixerConfig) (*MixedTrack, error) {
	t := config.Track

	module := &xmfile.Module{}

	// Start with a shallow copy.
	// We'll re-allocate the necessary slices later.
	*module = *t.Module

	mt := &MixedTrack{
		Track:  t,
		Module: module,
	}

	// +1 is needed for an empty pattern.
	module.NumPatterns = len(module.Patterns) + 1
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
				// n := module.Notes[noteID]
				// if n.Note == 0 || n.Note == 97 {
				// 	continue
				// }
				// if t.GetInstrumentKind(int(n.Instrument)) != edrum.UndefinedInstrument {
				// 	continue
				// }
				mixedRow.Notes[channelID] = noteID
			}
		}
	}
	// Fill the empty pattern.
	{
		numRows := 32
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

	stream := xm.NewStream()
	err := stream.LoadModule(module, xm.LoadModuleConfig{
		LinearInterpolation: true,
	})
	if err != nil {
		return nil, err
	}
	mt.Stream = stream

	return mt, nil
}
