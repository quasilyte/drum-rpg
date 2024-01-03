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
)
