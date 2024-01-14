package studio

import (
	"github.com/quasilyte/drum-rpg/assets"
	"github.com/quasilyte/drum-rpg/tracker"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/gmath"
)

type minimapNode struct {
	marker *ge.Sprite
	rects  []*ge.Rect
	track  *tracker.MixedTrack
}

func newMinimapNode(track *tracker.MixedTrack) *minimapNode {
	return &minimapNode{
		track: track,
	}
}

func (m *minimapNode) Init(scene *ge.Scene) {
	totalSize := m.totalHeight()
	numRows := float64(m.track.NumRows)

	m.rects = make([]*ge.Rect, len(m.track.Module.PatternOrder))
	offsetY := 16.0

	for i, patternIndex := range m.track.Module.PatternOrder {
		pat := &m.track.Module.Patterns[patternIndex]
		rectHeight := totalSize * (float64(len(pat.Rows)) / numRows)
		rect := ge.NewRect(scene.Context(), 24, rectHeight)
		rect.Centered = false
		if i == 0 {
			rect.FillColorScale.SetColor(ge.RGB(0xcdcdcd))
		} else {
			rect.FillColorScale.SetColor(getPatternColor(int(patternIndex)))
		}
		rect.OutlineColorScale.SetColor(ge.RGB(0x000000))
		rect.OutlineWidth = 1
		rect.Pos.Offset = gmath.Vec{
			X: 4,
			Y: (totalSize + 32) - rectHeight - offsetY,
		}
		offsetY += rectHeight
		m.rects[i] = rect
		scene.AddGraphics(rect)
	}

	m.marker = scene.NewSprite(assets.ImageMinimapMarker)
	m.marker.Pos.Offset.X = 16.0
	m.marker.Pos.Offset.Y = 120.0
	scene.AddGraphics(m.marker)
	m.UpdateMarker(0)
}

func (m *minimapNode) totalHeight() float64 {
	return (1080.0 / 2) - 32
}

func (m *minimapNode) UpdateMarker(t float64) {
	totalSize := m.totalHeight()
	progressPercent := t / m.track.Duration
	progressDist := totalSize * progressPercent
	y := totalSize - progressDist + 16.0
	m.marker.Pos.Offset.Y = y
}
