package scenes

import (
	"github.com/ebitenui/ebitenui/widget"
	"github.com/quasilyte/drum-rpg/assets"
	"github.com/quasilyte/drum-rpg/eui"
	"github.com/quasilyte/drum-rpg/session"
	"github.com/quasilyte/drum-rpg/styles"
	"github.com/quasilyte/ge"
)

type DrumkitListController struct {
	scene *ge.Scene
	state *session.State

	selectedDrumKitButton *widget.Button
	drumKitButtons        []*widget.Button

	createButton *widget.Button
	editButton   *widget.Button
	cloneButton  *widget.Button
}

func NewDrumkitListController(state *session.State) *DrumkitListController {
	return &DrumkitListController{
		state: state,
	}
}

func (c *DrumkitListController) Init(scene *ge.Scene) {
	c.scene = scene
	c.initUI()
}

func (c *DrumkitListController) selectKit(b *widget.Button) {
	if c.selectedDrumKitButton == b {
		return
	}
	if c.selectedDrumKitButton != nil {
		eui.ToggleButtonTextColor(c.state.UIResources, c.selectedDrumKitButton)
	}
	c.selectedDrumKitButton = b
	eui.ToggleButtonTextColor(c.state.UIResources, c.selectedDrumKitButton)
}

func (c *DrumkitListController) initUI() {
	root := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
		})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()))

	rowContainer := eui.NewRowLayoutContainerWithMinWidth(320, 8, nil)
	root.AddChild(rowContainer)

	rowContainer.AddChild(eui.NewCenteredLabel("Drum Kits", assets.BitmapFont3))

	{
		content := widget.NewContainer(
			widget.ContainerOpts.Layout(widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(16)),
			)),
		)

		drumKits := c.state.ContentDirManager.ListDrumKits()
		for _, k := range drumKits {
			var b *widget.Button
			b = eui.NewButtonWithConfig(c.state.UIResources, eui.ButtonConfig{
				ListStyle:  true,
				Text:       k,
				LayoutData: widget.RowLayoutData{Stretch: true},
				OnClick: func() {
					c.selectKit(b)
				},
			})
			content.AddChild(b)
			c.drumKitButtons = append(c.drumKitButtons, b)
		}

		if len(drumKits) != 0 {
			c.selectKit(c.drumKitButtons[0])
		}

		panel := eui.NewScrollContainer(c.state.UIResources, 196, content)
		rowContainer.AddChild(panel)
	}

	{
		grid := eui.NewGridContainer(3,
			widget.GridLayoutOpts.Spacing(8, 0),
			widget.GridLayoutOpts.Stretch([]bool{true, true, true}, nil))

		c.createButton = eui.NewButton(c.state.UIResources, "CREATE", func() {
		})

		c.editButton = eui.NewButton(c.state.UIResources, "EDIT", func() {
		})

		c.cloneButton = eui.NewButton(c.state.UIResources, "CLONE", func() {
		})

		grid.AddChild(c.createButton)

		grid.AddChild(c.editButton)

		grid.AddChild(c.cloneButton)

		rowContainer.AddChild(grid)
	}

	rowContainer.AddChild(eui.NewSeparator(nil, styles.TransparentColor))

	rowContainer.AddChild(eui.NewButton(c.state.UIResources, "BACK", func() {
		c.scene.Context().ChangeScene(NewStudioController(c.state))
	}))

	initUI(c.scene, root)
}

func (c *DrumkitListController) Update(delta float64) {}
