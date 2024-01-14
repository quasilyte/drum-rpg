package scenes

import (
	"github.com/ebitenui/ebitenui/widget"
	"github.com/quasilyte/drum-rpg/assets"
	"github.com/quasilyte/drum-rpg/eui"
	"github.com/quasilyte/drum-rpg/session"
	"github.com/quasilyte/ge"
)

type DrumkitEditController struct {
	scene *ge.Scene
	state *session.State
}

func NewDrumkitEditController(state *session.State) *DrumkitEditController {
	return &DrumkitEditController{
		state: state,
	}
}

func (c *DrumkitEditController) Init(scene *ge.Scene) {
	c.scene = scene
	c.initUI()
}

func (c *DrumkitEditController) initUI() {
	root := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
		})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()))

	rowContainer := eui.NewRowLayoutContainerWithMinWidth(320, 8, nil)
	root.AddChild(rowContainer)

	{
		grid := eui.NewGridContainer(3,
			widget.GridLayoutOpts.Stretch([]bool{true, true, true}, nil))
		// label := eui.NewLabel("Closed Hi-Hat", assets.BitmapFont1)
		grid.AddChild(eui.NewLabel("Closed Hi-Hat", assets.BitmapFont1))
		grid.AddChild(eui.NewLabel("Closed Hi-Hat", assets.BitmapFont1))
		grid.AddChild(eui.NewLabel("Closed Hi-Hat", assets.BitmapFont1))
		rowContainer.AddChild(grid)
	}

	{
		grid := eui.NewGridContainer(2,
			widget.GridLayoutOpts.Spacing(8, 0),
			widget.GridLayoutOpts.Stretch([]bool{true, true}, nil))

		grid.AddChild(eui.NewButton(c.state.UIResources, "BACK", func() {
			c.scene.Context().ChangeScene(NewStudioController(c.state))
		}))

		grid.AddChild(eui.NewButton(c.state.UIResources, "SAVE", func() {
		}))
		rowContainer.AddChild(grid)
	}

	initUI(c.scene, root)
}

func (c *DrumkitEditController) Update(delta float64) {}
