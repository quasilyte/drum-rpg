package assets

import (
	_ "image/png"

	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/ge"
)

func registerImageResources(ctx *ge.Context) {
	imageResources := map[resource.ImageID]resource.ImageInfo{
		ImageTrackBg: {Path: "image/studio/track_bg.png"},

		ImageNoteEffect: {Path: "image/studio/note_effect.png", FrameWidth: 32},

		ImageBarBass:        {Path: "image/studio/bass_bar.png"},
		ImageBarClosedHiHat: {Path: "image/studio/closed_hihat_bar.png"},
		ImageBarOpenHiHat:   {Path: "image/studio/open_hihat_bar.png"},
		ImageBarLeftTom:     {Path: "image/studio/left_tom_bar.png"},
		ImageBarRightTom:    {Path: "image/studio/right_tom_bar.png"},
	}

	for id, res := range imageResources {
		ctx.Loader.ImageRegistry.Set(id, res)
		ctx.Loader.LoadImage(id)
	}
}

const (
	ImageNone resource.ImageID = iota

	ImageTrackBg

	ImageNoteEffect

	ImageBarBass
	ImageBarClosedHiHat
	ImageBarOpenHiHat
	ImageBarLeftTom
	ImageBarRightTom
)
