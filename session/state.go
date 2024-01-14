package session

import (
	"github.com/quasilyte/drum-rpg/edrum"
	"github.com/quasilyte/drum-rpg/eui"
	"github.com/quasilyte/drum-rpg/midichan"
	"github.com/quasilyte/drum-rpg/tracker"
)

type State struct {
	MidiSystem *midichan.System
	SoundBanks []*edrum.SoundBank
	DrumKits   []*edrum.Kit
	Tracks     []*tracker.Track

	ContentDirManager *ContentDirManager

	UIResources *eui.Resources
}

func (s *State) FindSoundBank(name string) *edrum.SoundBank {
	for _, sb := range s.SoundBanks {
		if sb.Name == name {
			return sb
		}
	}
	return nil
}

func (s *State) FindDrumKit(name string) *edrum.Kit {
	for _, k := range s.DrumKits {
		if k.Name == name {
			return k
		}
	}
	return nil
}

func (s *State) FindTrack(name string) *tracker.Track {
	for _, t := range s.Tracks {
		if t.Name == name {
			return t
		}
	}
	return nil
}
