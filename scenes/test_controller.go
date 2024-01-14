package scenes

import (
	"github.com/quasilyte/drum-rpg/edrum"
	"github.com/quasilyte/drum-rpg/midichan"
	"github.com/quasilyte/drum-rpg/session"
	"github.com/quasilyte/drum-rpg/studio"
	"github.com/quasilyte/drum-rpg/tracker"
	"github.com/quasilyte/ge"
)

type TestController struct {
	state *session.State

	drumkit *edrum.Kit
	player  *midichan.Player

	runner *studio.Runner
}

func NewTestController(state *session.State) *TestController {
	return &TestController{state: state}
}

func (c *TestController) Init(scene *ge.Scene) {
	// track := c.state.FindTrack("HyperSpace")
	// track := c.state.FindTrack("Druuge Theme")
	track := c.state.FindTrack("Ilwrath Theme")
	// track := c.state.FindTrack("Thraddash Theme")
	// track := c.state.FindTrack("Arilou Theme")
	// track := c.state.FindTrack("Eye of The Tiger")
	// track := c.state.FindTrack("HyperSpace")
	if track == nil {
		panic("track not found")
	}
	mixedTrack, err := tracker.Mix(tracker.MixerConfig{
		Track: track,
	})
	if err != nil {
		panic(err)
	}

	trackPlayer, err := scene.Audio().GetContext().NewPlayer(mixedTrack.Stream)
	if err != nil {
		panic(err)
	}
	trackPlayer.Rewind()
	trackPlayer.SetVolume(0.65)
	trackPlayer.Play()

	// colombo := c.state.FindSoundBank("ColomboADK FreePats")
	// colombo := c.state.FindSoundBank("UQM Hyper")
	// colombo := c.state.FindSoundBank("UQM Arilou")
	// colombo := c.state.FindSoundBank("UQM Ilwrath")
	// c.drumkit = edrum.NewKit(edrum.RolandTD02kvMap, map[edrum.InstrumentKind]*edrum.Instrument{
	// 	edrum.BassInstrument:        edrum.NewInstrument(edrum.RandomSampleSelection, colombo.Samples[edrum.BassInstrument]),
	// 	edrum.LeftTomInstrument:     edrum.NewInstrument(edrum.RandomSampleSelection, colombo.Samples[edrum.LeftTomInstrument]),
	// 	edrum.RightTomInstrument:    edrum.NewInstrument(edrum.RandomSampleSelection, colombo.Samples[edrum.RightTomInstrument]),
	// 	edrum.FloorTomInstrument:    edrum.NewInstrument(edrum.RandomSampleSelection, colombo.Samples[edrum.FloorTomInstrument]),
	// 	edrum.SnareInstrument:       edrum.NewInstrument(edrum.RandomSampleSelection, colombo.Samples[edrum.SnareInstrument]),
	// 	edrum.LeftCymbalInstrument:  edrum.NewInstrument(edrum.RandomSampleSelection, colombo.Samples[edrum.LeftCymbalInstrument]),
	// 	edrum.RightCymbalInstrument: edrum.NewInstrument(edrum.RandomSampleSelection, colombo.Samples[edrum.RightCymbalInstrument]),
	// 	edrum.OpenHiHatInstrument:   edrum.NewInstrument(edrum.RandomSampleSelection, colombo.Samples[edrum.OpenHiHatInstrument]),
	// 	edrum.ClosedHiHatInstrument: edrum.NewInstrument(edrum.RandomSampleSelection, colombo.Samples[edrum.ClosedHiHatInstrument]),
	// })
	c.drumkit = c.state.FindDrumKit("Ilwrath")

	c.player = midichan.NewPlayer(c.drumkit)
	c.player.ConnectTo(c.state.MidiSystem)
	c.player.Init(scene)

	c.runner = studio.NewRunner(studio.RunnerConfig{
		State:      c.state,
		DrumPlayer: c.player,
		MixedTrack: mixedTrack,
		ErrorDist:  14,
	})
	c.runner.Init(scene)
}

func (c *TestController) Update(delta float64) {
	c.state.MidiSystem.Update()
	c.runner.Update(delta)
}
