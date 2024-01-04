package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/quasilyte/drum-rpg/assets"
	"github.com/quasilyte/drum-rpg/edrum"
	"github.com/quasilyte/drum-rpg/midichan"
	"github.com/quasilyte/drum-rpg/scenes"
	"github.com/quasilyte/drum-rpg/session"
	"github.com/quasilyte/drum-rpg/tracker"
	"github.com/quasilyte/ge"

	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

func main() {
	var contentDir string
	flag.StringVar(&contentDir, "content", "_content",
		`a path to a directory that contains the external game files`)
	flag.Parse()

	ctx := ge.NewContext(ge.ContextConfig{
		TimeDeltaMode: ge.TimeDeltaFixed120,
	})

	ctx.GameName = "drum_rpg"
	ctx.WindowTitle = "Drum RPG"
	ctx.WindowWidth = 1920.0 / 2
	ctx.WindowHeight = 1080.0 / 2
	ctx.FullScreen = true

	ctx.Loader.OpenAssetFunc = assets.MakeOpenAssetFunc(ctx)
	assets.RegisterResources(ctx)

	soundBanks, err := edrum.LoadSoundBanks(ctx, filepath.Join(contentDir, "soundbank"))
	if err != nil {
		panic(fmt.Sprintf("load sound banks: %v", err))
	}

	arilouTrack, err := tracker.ParseTrack("Star Control 2", filepath.Join(contentDir, "tracks", "uqm", "druuge.xm"))
	if err != nil {
		panic(err)
	}

	state := &session.State{
		MidiSystem: midichan.NewSystem(),
		SoundBanks: soundBanks,
		Tracks: []*tracker.Track{
			arilouTrack,
		},
	}

	if err := state.MidiSystem.Connect("TD-02:TD-02 MIDI"); err != nil {
		panic(err)
	}

	if err := ge.RunGame(ctx, scenes.NewTestController(state)); err != nil {
		panic(err)
	}
}
