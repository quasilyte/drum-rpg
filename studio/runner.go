package studio

import (
	"github.com/quasilyte/drum-rpg/assets"
	"github.com/quasilyte/drum-rpg/edrum"
	"github.com/quasilyte/drum-rpg/midichan"
	"github.com/quasilyte/drum-rpg/session"
	"github.com/quasilyte/drum-rpg/tracker"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
)

type Runner struct {
	scene *ge.Scene

	bg         *ge.Sprite
	state      *session.State
	drumPlayer *midichan.Player
	mixedTrack *tracker.MixedTrack

	t               float64
	visibleNotes    []*noteBarNode
	channelProgress [4]int
}

type RunnerConfig struct {
	State      *session.State
	DrumPlayer *midichan.Player
	MixedTrack *tracker.MixedTrack
}

func NewRunner(config RunnerConfig) *Runner {
	return &Runner{
		state:      config.State,
		drumPlayer: config.DrumPlayer,
		mixedTrack: config.MixedTrack,
	}
}

func (r *Runner) Init(scene *ge.Scene) {
	r.scene = scene

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
		channelID := info.Instrument.Channel()
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

func (r *Runner) Update(delta float64) {
	// Maybe spawn new notes.
	{
		for i, notes := range r.mixedTrack.InputChannels {
			currentNoteIndex := r.channelProgress[i]
			if currentNoteIndex >= len(notes) {
				continue // No more notes in this channel
			}
			currentNote := notes[currentNoteIndex]
			y := r.calcNoteY(currentNote.Time)
			if y < -32 {
				continue
			}
			r.channelProgress[i]++
			n := newNoteBar(i, currentNote)
			n.Pos.X = r.getChannelGatePos(i).X - 1
			n.Pos.Y = y
			n.Init(r.scene)
			r.visibleNotes = append(r.visibleNotes, n)
		}
	}

	// Advance the visible notes.
	{
		visibleNotes := r.visibleNotes[:0]
		for _, n := range r.visibleNotes {
			n.Pos.Y = r.calcNoteY(n.Time)
			if n.Pos.Y > (1080.0/2 + 32.0) {
				n.Dispose()
				continue
			}
			visibleNotes = append(visibleNotes, n)
		}
		r.visibleNotes = visibleNotes
	}

	r.t += delta
}

func (r *Runner) calcNoteY(noteTime float64) float64 {
	const (
		height           float64 = 1080.0 / 2
		gateThreshold    float64 = 64.0 / height
		secondsPerScreen float64 = 3.0
		gateTimeTail     float64 = secondsPerScreen * gateThreshold
	)
	timeBaseline := r.t - gateTimeTail + secondsPerScreen
	return ((timeBaseline - noteTime) / secondsPerScreen) * height
}

func (r *Runner) getChannelGatePos(id int) gmath.Vec {
	return gmath.Vec{
		X: r.bg.Pos.Offset.X + 49 + (float64(id) * 128),
		Y: 1080.0/2 - 64.0,
	}
}
