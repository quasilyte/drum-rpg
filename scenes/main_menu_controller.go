package scenes

import (
	"fmt"
	"os"
	"runtime"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/quasilyte/drum-rpg/assets"
	"github.com/quasilyte/drum-rpg/eui"
	"github.com/quasilyte/drum-rpg/gamedata"
	"github.com/quasilyte/drum-rpg/session"
	"github.com/quasilyte/drum-rpg/styles"
	"github.com/quasilyte/ge"
)

type MainMenuController struct {
	scene *ge.Scene
	state *session.State
}

func NewMainMenuController(state *session.State) *MainMenuController {
	return &MainMenuController{
		state: state,
	}
}

func (c *MainMenuController) Init(scene *ge.Scene) {
	c.scene = scene
	c.initUI()
}

func (c *MainMenuController) initUI() {
	root := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
		})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()))

	rowContainer := eui.NewRowLayoutContainerWithMinWidth(320, 8, nil)
	root.AddChild(rowContainer)

	rowContainer.AddChild(eui.NewCenteredLabel("XM Music Club", assets.BitmapFont3))

	rowContainer.AddChild(eui.NewSeparator(nil, styles.TransparentColor))

	rowContainer.AddChild(eui.NewButton(c.state.UIResources, "PLAY", func() {
		c.scene.Context().ChangeScene(NewPlayController(c.state))
	}))

	rowContainer.AddChild(eui.NewButton(c.state.UIResources, "SETTINGS", func() {
	}))

	b := eui.NewButton(c.state.UIResources, "CREDITS", func() {
		// TODO
	})
	b.GetWidget().Disabled = true
	rowContainer.AddChild(b)

	if runtime.GOARCH != "wasm" {
		rowContainer.AddChild(eui.NewButton(c.state.UIResources, "EXIT", func() {
			os.Exit(0)
		}))
	}

	rowContainer.AddChild(eui.NewSeparator(nil, styles.TransparentColor))
	rowContainer.AddChild(eui.NewCenteredLabel(fmt.Sprintf("Build %d", gamedata.CurrentBuild), assets.BitmapFont1))

	initUI(c.scene, root)
}

func (c *MainMenuController) Update(delta float64) {}
