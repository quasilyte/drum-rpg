package studio

import (
	"github.com/quasilyte/drum-rpg/assets"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
)

type floatingTextNode struct {
	pos        gmath.Vec
	text       string
	decaySpeed float64
	label      *ge.Label
}

func newFloatingTextNode(pos gmath.Vec, text string) *floatingTextNode {
	return &floatingTextNode{
		pos:  pos,
		text: text,
	}
}

func (t *floatingTextNode) IsDisposed() bool {
	return t.label.IsDisposed()
}

func (t *floatingTextNode) Init(scene *ge.Scene) {
	t.decaySpeed = scene.Rand().FloatRange(0.8, 1.2)
	t.label = ge.NewLabel(assets.BitmapFont1)
	t.label.Text = t.text
	t.label.Pos.Base = &t.pos
	t.label.Width = 128
	t.label.Height = 32
	t.label.Pos.Offset.X -= t.label.Width / 2
	t.label.Pos.Offset.Y -= t.label.Height / 2
	t.label.AlignHorizontal = ge.AlignHorizontalCenter
	t.label.AlignVertical = ge.AlignVerticalCenter
	t.label.SetAlpha(0.6)
	t.label.SetColorScaleRGBA(0x69, 0xe6, 0x69, 0xff)
	scene.AddGraphics(t.label)
}

func (t *floatingTextNode) Update(delta float64) {
	alpha := t.label.GetAlpha() - float32(delta*t.decaySpeed)
	if alpha < 0.05 {
		t.label.Dispose()
		return
	}
	t.label.SetAlpha(alpha)
	t.pos.Y -= 300 * delta
}
