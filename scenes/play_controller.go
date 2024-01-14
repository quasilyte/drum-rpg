package scenes

import (
	"github.com/ebitenui/ebitenui/widget"
	"github.com/quasilyte/drum-rpg/assets"
	"github.com/quasilyte/drum-rpg/eui"
	"github.com/quasilyte/drum-rpg/session"
	"github.com/quasilyte/drum-rpg/styles"
	"github.com/quasilyte/ge"
)

type PlayController struct {
	scene *ge.Scene
	state *session.State
}

func NewPlayController(state *session.State) *PlayController {
	return &PlayController{
		state: state,
	}
}

func (c *PlayController) Init(scene *ge.Scene) {
	c.scene = scene
	c.initUI()
}

func (c *PlayController) initUI() {
	root := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
		})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()))

	rowContainer := eui.NewRowLayoutContainerWithMinWidth(320, 8, nil)
	root.AddChild(rowContainer)

	rowContainer.AddChild(eui.NewCenteredLabel("Play", assets.BitmapFont3))

	{
		b := eui.NewButton(c.state.UIResources, "STORY MODE", func() {
		})
		rowContainer.AddChild(b)
		b.GetWidget().Disabled = true
	}

	{
		b := eui.NewButton(c.state.UIResources, "FREE MODE", func() {
		})
		rowContainer.AddChild(b)
		b.GetWidget().Disabled = true
	}

	{
		b := eui.NewButton(c.state.UIResources, "DEMO MODE", func() {
		})
		rowContainer.AddChild(b)
	}

	rowContainer.AddChild(eui.NewButton(c.state.UIResources, "STUDIO", func() {
		c.scene.Context().ChangeScene(NewStudioController(c.state))
	}))

	rowContainer.AddChild(eui.NewSeparator(nil, styles.TransparentColor))

	rowContainer.AddChild(eui.NewButton(c.state.UIResources, "BACK", func() {
		c.scene.Context().ChangeScene(NewMainMenuController(c.state))
	}))

	initUI(c.scene, root)
}

func (c *PlayController) Update(delta float64) {}
