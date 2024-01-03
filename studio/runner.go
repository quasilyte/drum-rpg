package studio

import (
	"fmt"

	"github.com/quasilyte/drum-rpg/assets"
	"github.com/quasilyte/drum-rpg/edrum"
	"github.com/quasilyte/drum-rpg/midichan"
	"github.com/quasilyte/drum-rpg/session"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
)

type Runner struct {
	bg         *ge.Sprite
	state      *session.State
	drumPlayer *midichan.Player
}

type RunnerConfig struct {
	State      *session.State
	DrumPlayer *midichan.Player
}

func NewRunner(config RunnerConfig) *Runner {
	return &Runner{
		state:      config.State,
		drumPlayer: config.DrumPlayer,
	}
}

func (r *Runner) Init(scene *ge.Scene) {
	bg := scene.NewSprite(assets.ImageTrackBg)
	bg.Centered = false
	bg.Pos.Offset.X = 32
	scene.AddGraphics(bg)
	r.bg = bg

	for i := 0; i < 4; i++ {
		pos := gmath.Vec{X: r.getChannelGatePos(i).X}
		l := newChannelBarNode(i, pos)
		scene.AddObject(l)
	}

	{
		beginPos := gmath.Vec{X: bg.Pos.Offset.X + 3, Y: 1080.0/2 - 64.0}
		endPos := gmath.Vec{X: bg.Pos.Offset.X + bg.ImageWidth() - 3, Y: 1080.0/2 - 64.0}
		triggerLine := ge.NewLine(ge.Pos{Offset: beginPos}, ge.Pos{Offset: endPos})
		triggerLine.Width = 1
		triggerLine.SetColorScaleRGBA(0x69, 0xe6, 0x69, 0xff)
		scene.AddGraphics(triggerLine)
	}

	r.drumPlayer.EventNote.Connect(nil, func(info midichan.NoteEventInfo) {
		channelID := r.getInstrumentChannel(info.Instrument)
		pos := r.getChannelGatePos(channelID)

		e := newEffectNode(pos.Add(gmath.Vec{X: scene.Rand().FloatRange(-2, 2)}), assets.ImageNoteEffect)
		e.noFlip = true
		scene.AddObject(e)
		e.anim.SetSecondsPerFrame(0.03)

		instrumentLabel := "?"
		switch info.Instrument {
		case edrum.BassInstrument:
			instrumentLabel = "kick!"
		case edrum.SnareInstrument:
			instrumentLabel = "snare!"
		case edrum.OpenHiHatInstrument:
			instrumentLabel = "open hi-hat!"
		case edrum.ClosedHiHatInstrument:
			instrumentLabel = "closed hi-hat!"
		case edrum.LeftCymbalInstrument:
			instrumentLabel = "crash!"
		case edrum.LeftTomInstrument:
			instrumentLabel = "high tom!"
		case edrum.RightTomInstrument:
			instrumentLabel = "mid tom!"
		case edrum.FloorTomInstrument:
			instrumentLabel = "floor tom!"
		}

		textPos := pos.Sub(gmath.Vec{Y: 14}).Add(gmath.Vec{X: scene.Rand().FloatRange(-4, 4)})
		t := newFloatingTextNode(textPos, instrumentLabel)
		scene.AddObject(t)
	})
}

func (r *Runner) Update(delta float64) {}

func (r *Runner) getChannelGatePos(id int) gmath.Vec {
	return gmath.Vec{
		X: r.bg.Pos.Offset.X + 49 + (float64(id) * 128),
		Y: 1080.0/2 - 64.0,
	}
}

func (r *Runner) getInstrumentChannel(kind edrum.InstrumentKind) int {
	switch kind {
	case edrum.OpenHiHatInstrument, edrum.ClosedHiHatInstrument, edrum.LeftCymbalInstrument, edrum.RightCymbalInstrument:
		return 0
	case edrum.SnareInstrument:
		return 1
	case edrum.LeftTomInstrument, edrum.RightTomInstrument, edrum.FloorTomInstrument:
		return 2
	case edrum.BassInstrument:
		return 3
	}

	panic(fmt.Errorf("unexpected instrument kind: %v", kind))
}
