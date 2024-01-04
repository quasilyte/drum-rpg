package midichan

import (
	"github.com/quasilyte/drum-rpg/edrum"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gsignal"
	"gitlab.com/gomidi/midi/v2"
)

type Player struct {
	drumKit *edrum.Kit
	scene   *ge.Scene

	EventNote gsignal.Event[NoteEventInfo]
}

type NoteEventInfo struct {
	Instrument edrum.InstrumentKind
	Volume     float64
}

func NewPlayer(drumKit *edrum.Kit) *Player {
	return &Player{
		drumKit: drumKit,
	}
}

func (p *Player) Init(scene *ge.Scene) {
	p.scene = scene
}

func (p *Player) ConnectTo(sys *System) {
	sys.OnMessage = p.onMidiMessage
}

func (p *Player) calculateVolume(velocity uint8) float64 {
	return float64(velocity) * (1.0 / 127.0)
}

func (p *Player) onMidiMessage(msg midi.Message) {
	typ := msg.Type()

	switch typ {
	case midi.NoteOnMsg:
		var channel uint8
		var note uint8
		var velocity uint8
		msg.GetNoteOn(&channel, &note, &velocity)
		instKind := p.drumKit.NoteMap[note]
		if instKind == edrum.UndefinedInstrument {
			// fmt.Println("note on: undefined instrument")
			return
		}
		inst := p.drumKit.InstrumentMap[instKind]
		if inst == nil {
			// fmt.Println("note on: nil instrument for a note")
			return
		}
		samples := inst.GetSamples(velocity)
		if len(samples) == 0 {
			// fmt.Println("note on: found no samples for the note", velocity, note, instKind)
			return
		}
		sampleID := gmath.RandElem(p.scene.Rand(), samples)
		vol := p.calculateVolume(velocity)
		p.EventNote.Emit(NoteEventInfo{
			Instrument: instKind,
			Volume:     vol,
		})
		// fmt.Println("playing sample with id", sampleID, "and vol", vol, "velocity is", velocity)
		p.scene.Audio().PlaySoundWithVolume(sampleID, vol)
	}
}
