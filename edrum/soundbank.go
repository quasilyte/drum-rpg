package edrum

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/ge"
)

type SoundBank struct {
	Name string

	Samples map[InstrumentKind][]Sample

	NumSamples int
}

func LoadSoundBanks(ctx *ge.Context, dir string) ([]*SoundBank, error) {
	list, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	audioID := resource.AudioID(100)

	var result []*SoundBank
	for _, f := range list {
		if !f.IsDir() {
			continue
		}
		name := f.Name()
		sb, err := loadSoundBank(ctx, audioID, dir, name)
		if err != nil {
			return nil, fmt.Errorf("load %q: %w", name, err)
		}
		audioID += resource.AudioID(sb.NumSamples)
		result = append(result, sb)
	}

	return result, nil
}

func loadSoundBank(ctx *ge.Context, audioID resource.AudioID, dir, name string) (*SoundBank, error) {
	configData, err := os.ReadFile(filepath.Join(dir, name, "config.json"))
	if err != nil {
		return nil, err
	}

	sb := &SoundBank{
		Name:    name,
		Samples: make(map[InstrumentKind][]Sample),
	}
	id := audioID

	type sampleInfo struct {
		Name        string
		MinVelocity float64
		MaxVelocity float64
	}
	var soundBankConfig map[string][]sampleInfo
	if err := json.Unmarshal(configData, &soundBankConfig); err != nil {
		return nil, err
	}

	for instrumentTag, samples := range soundBankConfig {
		kind := InstrumentKindByName(instrumentTag)
		if kind == UndefinedInstrument {
			return nil, fmt.Errorf("unrecognized instrument type: %q", instrumentTag)
		}
		for _, info := range samples {
			filename := info.Name + ".wav"
			fullPath := filepath.Join(dir, name, filename)
			if !fileExists(fullPath) {
				return nil, fmt.Errorf("sample %q can't be located: file does not exist", info.Name)
			}
			ctx.Loader.AudioRegistry.Set(id, resource.AudioInfo{
				Path: "$soundbank/" + fullPath,
			})
			fmt.Println("load", "$soundbank/"+fullPath)
			ctx.Loader.LoadAudio(id)

			s := Sample{
				Name:        info.Name,
				MinVelocity: info.MinVelocity,
				MaxVelocity: info.MaxVelocity,
				AudioID:     id,
			}
			sb.Samples[kind] = append(sb.Samples[kind], s)
			sb.NumSamples++

			id++
		}
	}

	return sb, nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
