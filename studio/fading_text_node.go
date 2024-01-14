package studio

import (
	"github.com/quasilyte/drum-rpg/assets"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
)

type fadingTextNode struct {
	pos   gmath.Vec
	text  string
	delay float64
	label *ge.Label
}

func newFadingTextNode(pos gmath.Vec, text string) *fadingTextNode {
	return &fadingTextNode{
		pos:  pos,
		text: text,
	}
}

func (t *fadingTextNode) IsDisposed() bool {
	return t.label.IsDisposed()
}

func (t *fadingTextNode) Init(scene *ge.Scene) {
	t.delay = 3.0
	t.label = ge.NewLabel(assets.BitmapFont4)
	t.label.Text = t.text
	t.label.Pos.Base = &t.pos
	t.label.Width = 400
	t.label.Height = 96
	t.label.Pos.Offset.X -= t.label.Width / 2
	t.label.Pos.Offset.Y -= t.label.Height / 2
	t.label.AlignHorizontal = ge.AlignHorizontalCenter
	t.label.AlignVertical = ge.AlignVerticalCenter
	t.label.SetColorScaleRGBA(0x69, 0xe6, 0x69, 0xff)
	scene.AddGraphics(t.label)
}

func (t *fadingTextNode) Update(delta float64) {
	if t.delay > 0 {
		t.delay -= delta
		return
	}

	alpha := t.label.GetAlpha() - float32(delta*0.5)
	if alpha < 0.05 {
		t.label.Dispose()
		return
	}
	t.label.SetAlpha(alpha)
}
