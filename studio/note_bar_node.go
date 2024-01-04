package studio

import (
	"github.com/quasilyte/drum-rpg/assets"
	"github.com/quasilyte/drum-rpg/edrum"
	"github.com/quasilyte/drum-rpg/tracker"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
)

type noteBarNode struct {
	Time       float64
	Instrument edrum.InstrumentKind
	Pos        gmath.Vec
	ChannelID  int

	sprite *ge.Sprite
}

func newNoteBar(channelID int, n tracker.InputNote) *noteBarNode {
	return &noteBarNode{
		ChannelID:  channelID,
		Time:       n.Time,
		Instrument: n.Instrument,
	}
}

func (b *noteBarNode) Dispose() {
	b.sprite.Dispose()
}

func (b *noteBarNode) Init(scene *ge.Scene) {
	var img resource.ImageID
	switch b.Instrument {
	case edrum.BassInstrument, edrum.SnareInstrument:
		img = assets.ImageBarBass
	case edrum.ClosedHiHatInstrument:
		img = assets.ImageBarClosedHiHat
	case edrum.OpenHiHatInstrument:
		img = assets.ImageBarOpenHiHat
	case edrum.LeftTomInstrument:
		img = assets.ImageBarLeftTom
	case edrum.RightTomInstrument:
		img = assets.ImageBarRightTom
	}
	b.sprite = scene.NewSprite(img)
	b.sprite.Pos.Base = &b.Pos
	scene.AddGraphics(b.sprite)
}
