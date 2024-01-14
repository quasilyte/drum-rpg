package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/quasilyte/drum-rpg/assets"
	"github.com/quasilyte/drum-rpg/edrum"
	"github.com/quasilyte/drum-rpg/eui"
	"github.com/quasilyte/drum-rpg/fileutil"
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

	log.SetFlags(0)

	{
		// Prefer an absolute path, whether possible.
		absContentDir, err := filepath.Abs(contentDir)
		if err == nil {
			contentDir = absContentDir
		}
	}
	if !fileutil.FileExists(contentDir) {
		log.Fatalf("specified --content directory (%q) doesn't exist", contentDir)
	}

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

	arilouTrack, err := tracker.ParseTrack("Star Control 2", filepath.Join(contentDir, "tracks", "uqm", "ilwrath.xm"))
	// arilouTrack, err := tracker.ParseTrack("Star Control 2", filepath.Join(contentDir, "tracks", "extra", "eye_of_the_tiger.xm"))
	if err != nil {
		panic(err)
	}

	state := &session.State{
		MidiSystem: midichan.NewSystem(),
		SoundBanks: soundBanks,
		Tracks: []*tracker.Track{
			arilouTrack,
		},
		ContentDirManager: &session.ContentDirManager{
			Dir: contentDir,
		},
		UIResources: eui.PrepareResources(ctx.Loader),
	}

	if err := state.ContentDirManager.Init(); err != nil {
		panic(err)
	}

	{
		kits, err := session.LoadDrumKits(state, filepath.Join(contentDir, "drumkits"))
		if err != nil {
			panic(err)
		}
		state.DrumKits = kits
	}

	// if err := state.MidiSystem.Connect("TD-02:TD-02 MIDI"); err != nil {
	// 	panic(err)
	// }

	// if err := ge.RunGame(ctx, scenes.NewTestController(state)); err != nil {
	if err := ge.RunGame(ctx, scenes.NewMainMenuController(state)); err != nil {
		panic(err)
	}
}
