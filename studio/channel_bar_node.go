package studio

import (
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
)

type channelBarNode struct {
	id   int
	pos  gmath.Vec
	line *ge.Line
}

func newChannelBarNode(id int, pos gmath.Vec) *channelBarNode {
	return &channelBarNode{
		id:  id,
		pos: pos,
	}
}

func (b *channelBarNode) Init(scene *ge.Scene) {
	lineBegin := ge.Pos{Offset: b.pos}
	lineEnd := ge.Pos{Offset: b.pos.Add(gmath.Vec{Y: 1080.0 / 2})}
	b.line = ge.NewLine(lineBegin, lineEnd)
	b.line.Width = 2
	b.line.SetColorScaleRGBA(0x69, 0xe6, 0x69, 0xff/3)
	scene.AddGraphics(b.line)
}

func (b *channelBarNode) IsDisposed() bool { return false }

func (b *channelBarNode) Update(delta float64) {}
