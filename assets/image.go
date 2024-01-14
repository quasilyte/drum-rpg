package assets

import (
	_ "image/png"

	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/ge"
)

func registerImageResources(ctx *ge.Context) {
	imageResources := map[resource.ImageID]resource.ImageInfo{
		ImageUIButtonIdle:          {Path: "image/ebitenui/button-idle.png"},
		ImageUIButtonDisabled:      {Path: "image/ebitenui/button-disabled.png"},
		ImageUIButtonHover:         {Path: "image/ebitenui/button-hover.png"},
		ImageUIButtonPressed:       {Path: "image/ebitenui/button-pressed.png"},
		ImageUIListButtonIdle:      {Path: "image/ebitenui/list-button-idle.png"},
		ImageUIListButtonDisabled:  {Path: "image/ebitenui/list-button-disabled.png"},
		ImageUIListButtonHover:     {Path: "image/ebitenui/list-button-hover.png"},
		ImageUIListButtonPressed:   {Path: "image/ebitenui/list-button-pressed.png"},
		ImageUIPanel:               {Path: "image/ebitenui/panel.png"},
		ImageUIListPanelIdle:       {Path: "image/ebitenui/list-panel-idle.png"},
		ImageUIListPanelMask:       {Path: "image/ebitenui/list-panel-mask.png"},
		ImageUISliderHandleIdle:    {Path: "image/ebitenui/slider-handle-idle.png"},
		ImageUISliderHandleHover:   {Path: "image/ebitenui/slider-handle-hover.png"},
		ImageUISliderHandlePressed: {Path: "image/ebitenui/slider-handle-pressed.png"},

		ImageMenuBg:  {Path: "image/menu_bg.png"},
		ImageTrackBg: {Path: "image/studio/track_bg.png"},

		ImageNoteEffect:    {Path: "image/studio/note_effect.png", FrameWidth: 32},
		ImageNoteHitEffect: {Path: "image/studio/note_hit_effect.png", FrameWidth: 64},

		ImageMinimapMarker: {Path: "image/studio/minimap_marker.png"},

		ImageBarBass:        {Path: "image/studio/bass_bar.png"},
		ImageBarLeftCymbal:  {Path: "image/studio/crash_bar.png"},
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

	ImageUIButtonIdle
	ImageUIButtonDisabled
	ImageUIButtonHover
	ImageUIButtonPressed
	ImageUIListButtonIdle
	ImageUIListButtonDisabled
	ImageUIListButtonHover
	ImageUIListButtonPressed
	ImageUIPanel
	ImageUIListPanelIdle
	ImageUIListPanelMask
	ImageUISliderHandleIdle
	ImageUISliderHandleHover
	ImageUISliderHandlePressed

	ImageMenuBg
	ImageTrackBg

	ImageNoteEffect
	ImageNoteHitEffect

	ImageMinimapMarker

	ImageBarBass
	ImageBarClosedHiHat
	ImageBarOpenHiHat
	ImageBarLeftTom
	ImageBarRightTom
	ImageBarLeftCymbal
)
