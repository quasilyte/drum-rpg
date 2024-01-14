package scenes

import (
	"github.com/ebitenui/ebitenui/widget"
	"github.com/quasilyte/drum-rpg/assets"
	"github.com/quasilyte/drum-rpg/eui"
	"github.com/quasilyte/drum-rpg/session"
	"github.com/quasilyte/drum-rpg/styles"
	"github.com/quasilyte/ge"
)

type StudioController struct {
	scene *ge.Scene
	state *session.State
}

func NewStudioController(state *session.State) *StudioController {
	return &StudioController{
		state: state,
	}
}

func (c *StudioController) Init(scene *ge.Scene) {
	c.scene = scene
	c.initUI()
}

func (c *StudioController) initUI() {
	root := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
		})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()))

	rowContainer := eui.NewRowLayoutContainerWithMinWidth(320, 8, nil)
	root.AddChild(rowContainer)

	rowContainer.AddChild(eui.NewCenteredLabel("Studio", assets.BitmapFont3))

	rowContainer.AddChild(eui.NewButton(c.state.UIResources, "EDIT DRUM KITS", func() {
		c.scene.Context().ChangeScene(NewDrumkitListController(c.state))
	}))

	rowContainer.AddChild(eui.NewSeparator(nil, styles.TransparentColor))

	rowContainer.AddChild(eui.NewButton(c.state.UIResources, "BACK", func() {
		c.scene.Context().ChangeScene(NewPlayController(c.state))
	}))

	initUI(c.scene, root)
}

func (c *StudioController) Update(delta float64) {}
