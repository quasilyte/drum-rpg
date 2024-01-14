package session

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/quasilyte/drum-rpg/contentdir"
	"github.com/quasilyte/drum-rpg/edrum"
	"github.com/quasilyte/drum-rpg/jsonc"
)

func LoadDrumKits(state *State, dir string) ([]*edrum.Kit, error) {
	list, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var result []*edrum.Kit
	for _, f := range list {
		if !f.Type().IsRegular() {
			continue
		}
		name := f.Name()
		k, err := loadDrumKit(state, filepath.Join(dir, name))
		if err != nil {
			return nil, fmt.Errorf("load %q kit: %w", name, err)
		}
		result = append(result, k)
	}
	return result, nil
}

func loadDrumKit(state *State, filename string) (*edrum.Kit, error) {
	configData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var drumKitConfig contentdir.DrumKit
	if err := jsonc.Unmarshal(configData, &drumKitConfig); err != nil {
		return nil, err
	}

	instruments := make(map[edrum.InstrumentKind]*edrum.Instrument)
	for _, info := range drumKitConfig.Instruments {
		keyParts := strings.Split(info.Key, "/")
		if len(keyParts) != 3 {
			return nil, fmt.Errorf("invalid key %q", info.Key)
		}
		sb := state.FindSoundBank(keyParts[0])
		if sb == nil {
			return nil, fmt.Errorf("%q: can't find referenced soundbank %q", info.Key, keyParts[0])
		}
		instKind := edrum.InstrumentKindByName(keyParts[1])
		if instKind == edrum.UndefinedInstrument {
			return nil, fmt.Errorf("%q: unknown instrument kind: %q", info.Key, keyParts[1])
		}
		var samples []edrum.Sample
		for _, s := range sb.Samples[instKind] {
			if s.Tag != keyParts[2] {
				continue
			}
			samples = append(samples, s)
		}
		inst := edrum.NewInstrument(edrum.RandomSampleSelection, samples)

		instruments[instKind] = inst
	}

	k := edrum.NewKit(edrum.RolandTD02kvMap, instruments)
	k.Name = drumKitConfig.Name
	return k, nil
}
