package session

import (
	"github.com/quasilyte/drum-rpg/edrum"
	"github.com/quasilyte/drum-rpg/midichan"
)

type State struct {
	MidiSystem *midichan.System
	SoundBanks []*edrum.SoundBank
}

func (s *State) FindSoundBank(name string) *edrum.SoundBank {
	for _, sb := range s.SoundBanks {
		if sb.Name == name {
			return sb
		}
	}
	return nil
}
