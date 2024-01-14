package eui

import (
	"image/color"
	"math"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/drum-rpg/assets"
	"github.com/quasilyte/drum-rpg/styles"
	resource "github.com/quasilyte/ebitengine-resource"
	"golang.org/x/image/font"
)

type Resources struct {
	loader     *resource.Loader
	button     *buttonResource
	listButton *buttonResource
	panel      *panelResource
	list       *listResource
}

type listResource struct {
	Image   *widget.ScrollContainerImage
	Handle  *widget.ButtonImage
	Padding widget.Insets
}

type buttonResource struct {
	Image         *widget.ButtonImage
	Padding       widget.Insets
	TextColors    *widget.ButtonTextColor
	AltTextColors *widget.ButtonTextColor
	FontFace      font.Face
}

type panelResource struct {
	Image   *image.NineSlice
	Padding widget.Insets
}

func PrepareResources(loader *resource.Loader) *Resources {
	result := &Resources{
		loader: loader,
	}

	smallFont := assets.BitmapFont1
	normalFont := assets.BitmapFont2

	{
		disabled := nineSliceImage(loader.LoadImage(assets.ImageUIButtonDisabled).Data, 12, 30)
		idle := nineSliceImage(loader.LoadImage(assets.ImageUIButtonIdle).Data, 12, 30)
		hover := nineSliceImage(loader.LoadImage(assets.ImageUIButtonHover).Data, 12, 30)
		pressed := nineSliceImage(loader.LoadImage(assets.ImageUIButtonPressed).Data, 12, 30)
		buttonPadding := widget.Insets{
			Left:   24,
			Right:  24,
			Top:    4,
			Bottom: 4,
		}
		buttonColors := &widget.ButtonTextColor{
			Idle:     styles.NormalTextColor,
			Disabled: styles.DisabledTextColor,
		}
		result.button = &buttonResource{
			Image: &widget.ButtonImage{
				Idle:     idle,
				Hover:    hover,
				Pressed:  pressed,
				Disabled: disabled,
			},
			Padding:    buttonPadding,
			TextColors: buttonColors,
			AltTextColors: &widget.ButtonTextColor{
				Idle:     styles.NormalTextColor,
				Disabled: styles.DisabledTextColor,
			},
			FontFace: normalFont,
		}
	}

	{
		disabled := nineSliceImage(loader.LoadImage(assets.ImageUIListButtonDisabled).Data, 12, 30)
		idle := nineSliceImage(loader.LoadImage(assets.ImageUIListButtonIdle).Data, 12, 30)
		hover := nineSliceImage(loader.LoadImage(assets.ImageUIListButtonHover).Data, 12, 30)
		pressed := nineSliceImage(loader.LoadImage(assets.ImageUIListButtonPressed).Data, 12, 30)
		buttonPadding := widget.Insets{
			Left:   16,
			Right:  16,
			Top:    4,
			Bottom: 4,
		}
		buttonColors := &widget.ButtonTextColor{
			Idle:     styles.NormalTextColor,
			Disabled: styles.DisabledTextColor,
		}
		result.listButton = &buttonResource{
			Image: &widget.ButtonImage{
				Idle:     idle,
				Hover:    hover,
				Pressed:  pressed,
				Disabled: disabled,
			},
			Padding:    buttonPadding,
			TextColors: buttonColors,
			AltTextColors: &widget.ButtonTextColor{
				Idle:     styles.SelectedTextColor,
				Disabled: styles.DisabledTextColor,
			},
			FontFace: smallFont,
		}
	}

	{
		idle := loader.LoadImage(assets.ImageUIPanel).Data
		result.panel = &panelResource{
			Image: nineSliceImage(idle, 10, 10),
			Padding: widget.Insets{
				Left:   16,
				Right:  16,
				Top:    10,
				Bottom: 10,
			},
		}
	}

	{
		idle := loader.LoadImage(assets.ImageUIListPanelIdle).Data
		result.list = &listResource{
			Image: &widget.ScrollContainerImage{
				Idle: nineSliceImage(idle, 10, 10),
				// Disabled: image.NewNineSlice(disabled, [3]int{25, 12, 22}, [3]int{25, 12, 25}),
				Mask: nineSliceImage(loader.LoadImage(assets.ImageUIListPanelMask).Data, 10, 10),
			},
			Handle: &widget.ButtonImage{
				Idle:    nineSliceImage(loader.LoadImage(assets.ImageUISliderHandleIdle).Data, 5, 10),
				Hover:   nineSliceImage(loader.LoadImage(assets.ImageUISliderHandleHover).Data, 5, 10),
				Pressed: nineSliceImage(loader.LoadImage(assets.ImageUISliderHandlePressed).Data, 5, 10),
			},
		}
	}

	return result
}

func nineSliceImage(i *ebiten.Image, centerWidth, centerHeight int) *image.NineSlice {
	w, h := i.Size()
	return image.NewNineSlice(i,
		[3]int{(w - centerWidth) / 2, centerWidth, w - (w-centerWidth)/2 - centerWidth},
		[3]int{(h - centerHeight) / 2, centerHeight, h - (h-centerHeight)/2 - centerHeight})
}

func NewGraphic(res *Resources, imageID resource.ImageID) *widget.Graphic {
	img := res.loader.LoadImage(imageID)
	return widget.NewGraphic(
		widget.GraphicOpts.Image(img.Data),
		widget.GraphicOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
	)
}

type ButtonConfig struct {
	Text         string
	TextAltColor bool
	ListStyle    bool
	OnClick      func()
	LayoutData   any
	MinWidth     int
	Font         font.Face
	AlignLeft    bool
}

func NewButtonWithConfig(res *Resources, config ButtonConfig) *widget.Button {
	buttonRes := res.button
	if config.ListStyle {
		buttonRes = res.listButton
	}

	ff := config.Font
	if ff == nil {
		ff = buttonRes.FontFace
	}
	options := []widget.ButtonOpt{
		widget.ButtonOpts.Image(buttonRes.Image),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			if config.OnClick != nil {
				config.OnClick()
			}
		}),
	}
	colors := buttonRes.TextColors
	if config.TextAltColor {
		colors = buttonRes.AltTextColors
	}
	options = append(options,
		widget.ButtonOpts.Text(config.Text, ff, colors),
		widget.ButtonOpts.TextPadding(buttonRes.Padding))
	if config.AlignLeft {
		options = append(options, widget.ButtonOpts.TextPosition(widget.TextPositionStart, widget.TextPositionCenter))
	}
	if config.LayoutData != nil {
		options = append(options, widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(config.LayoutData)))
	}
	if config.MinWidth != 0 {
		options = append(options, widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.MinSize(config.MinWidth, 0)))
	}
	return widget.NewButton(options...)
}

func ToggleButtonTextColor(res *Resources, b *widget.Button) {
	switch b.TextColor {
	case res.listButton.TextColors:
		b.TextColor = res.listButton.AltTextColors
	case res.listButton.AltTextColors:
		b.TextColor = res.listButton.TextColors
	}
}

func NewButton(res *Resources, text string, onclick func()) *widget.Button {
	return NewButtonWithConfig(res, ButtonConfig{
		Text:    text,
		OnClick: onclick,
	})
}

func NewCenteredLabel(text string, ff font.Face) *widget.Text {
	return NewCenteredLabelWithMaxWidth(text, ff, -1)
}

func NewCenteredLabelWithMaxWidth(text string, ff font.Face, width float64) *widget.Text {
	options := []widget.TextOpt{
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			}),
		),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.Text(text, ff, styles.NormalTextColor),
	}
	if width != -1 {
		options = append(options, widget.TextOpts.MaxWidth(width))
	}
	return widget.NewText(options...)
}

func NewColoredLabel(text string, ff font.Face, clr color.RGBA, options ...widget.TextOpt) *widget.Text {
	opts := []widget.TextOpt{
		widget.TextOpts.Text(text, ff, clr),
	}
	if len(options) != 0 {
		opts = append(opts, options...)
	}
	return widget.NewText(opts...)
}

func NewLabel(text string, ff font.Face, options ...widget.TextOpt) *widget.Text {
	return NewColoredLabel(text, ff, styles.NormalTextColor, options...)
}

func NewSeparator(ld interface{}, clr color.RGBA) widget.PreferredSizeLocateableWidget {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top:    20,
				Bottom: 20,
			}))),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(ld)))

	c.AddChild(widget.NewGraphic(
		widget.GraphicOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch:   true,
			MaxHeight: 2,
		})),
		widget.GraphicOpts.ImageNineSlice(image.NewNineSliceColor(clr)),
	))

	return c
}

func NewPanelWithPadding(res *Resources, minWidth, minHeight int, padding widget.Insets) *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(res.panel.Image),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(padding),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				StretchHorizontal:  true,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
			widget.WidgetOpts.MinSize(minWidth, minHeight),
		),
	)
}

func NewScrollContainer(res *Resources, maxHeight int, content *widget.Container, options ...widget.ListOpt) *widget.Container {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Spacing(2, 0),
			widget.GridLayoutOpts.Stretch([]bool{true, false}, []bool{true}),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.GridLayoutData{
				MaxHeight: maxHeight,
			}),
		),
	)

	scrollContainer := widget.NewScrollContainer(
		widget.ScrollContainerOpts.Content(content),
		widget.ScrollContainerOpts.StretchContentWidth(),
		widget.ScrollContainerOpts.Image(res.list.Image),
	)
	c.AddChild(scrollContainer)

	pageSizeFunc := func() int {
		return int(math.Round(float64(scrollContainer.ContentRect().Dy()) / float64(content.GetWidget().Rect.Dy()) * 1000))
	}

	vSlider := widget.NewSlider(
		widget.SliderOpts.Direction(widget.DirectionVertical),
		widget.SliderOpts.MinMax(0, 1000),
		widget.SliderOpts.PageSizeFunc(pageSizeFunc),
		widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
			scrollContainer.ScrollTop = float64(args.Slider.Current) / 1000
		}),
		widget.SliderOpts.Images(
			&widget.SliderTrackImage{
				Idle:  image.NewNineSliceColor(color.NRGBA{}),
				Hover: image.NewNineSliceColor(color.NRGBA{}),
			},
			&widget.ButtonImage{
				Idle:    res.list.Handle.Idle,
				Hover:   res.list.Handle.Hover,
				Pressed: res.list.Handle.Pressed,
			},
		),
	)
	c.AddChild(vSlider)

	// Set the slider's position if the scrollContainer is scrolled by other means than the slider.
	scrollContainer.GetWidget().ScrolledEvent.AddHandler(func(args interface{}) {
		a := args.(*widget.WidgetScrolledEventArgs)
		p := pageSizeFunc() / 3
		if p < 1 {
			p = 1
		}
		vSlider.Current -= int(math.Round(a.Y * float64(p)))
	})

	return c
}
