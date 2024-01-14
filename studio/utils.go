package studio

import (
	"image/color"

	"github.com/quasilyte/ge"
)

func getPatternColor(patternIndex int) color.RGBA {
	switch patternIndex % 4 {
	case 1:
		return ge.RGB(0x698ad2)
	case 2:
		return ge.RGB(0xc78ac7)
	case 3:
		return ge.RGB(0xc24554)
	default:
		return ge.RGB(0x69e669)
	}
}

func getPatternColorScale(patternIndex int) ge.ColorScale {
	switch patternIndex % 4 {
	case 1:
		// Blue color.
		return ge.ColorScale{R: 1.0, G: 0.6, B: 2.0, A: 1}
	case 2:
		// Pink color.
		return ge.ColorScale{R: 1.9, G: 0.6, B: 1.9, A: 1}
	case 3:
		// Red color.
		return ge.ColorScale{R: 1.85, G: 0.3, B: 0.8, A: 1}
	default:
		// Normal (green) color.
		return ge.ColorScale{R: 1, G: 1, B: 1, A: 1}
	}
}
