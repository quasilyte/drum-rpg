package edrum

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/quasilyte/drum-rpg/fileutil"
	"github.com/quasilyte/drum-rpg/jsonc"
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
		sb, err := loadSoundBank(ctx, audioID, filepath.Join(dir, name))
		if err != nil {
			return nil, fmt.Errorf("load %q soundbank: %w", name, err)
		}
		audioID += resource.AudioID(sb.NumSamples)
		result = append(result, sb)
	}

	return result, nil
}

func loadSoundBank(ctx *ge.Context, audioID resource.AudioID, dir string) (*SoundBank, error) {
	configData, err := os.ReadFile(filepath.Join(dir, "config.jsonc"))
	if err != nil {
		return nil, err
	}

	sb := &SoundBank{
		Samples: make(map[InstrumentKind][]Sample),
	}
	id := audioID

	type sampleInfo struct {
		File        string
		MinVelocity float64
		MaxVelocity float64
	}
	type instrumentInfo struct {
		Kind    string
		Tag     string
		Samples []sampleInfo
	}
	var soundBankConfig struct {
		Name        string
		Instruments []instrumentInfo
	}
	if err := jsonc.Unmarshal(configData, &soundBankConfig); err != nil {
		return nil, err
	}

	sb.Name = soundBankConfig.Name

	for _, instrument := range soundBankConfig.Instruments {
		kind := InstrumentKindByName(instrument.Kind)
		if kind == UndefinedInstrument {
			return nil, fmt.Errorf("unrecognized instrument type: %q", instrument.Kind)
		}
		for _, info := range instrument.Samples {
			if !strings.HasSuffix(info.File, ".wav") {
				return nil, fmt.Errorf("sample %q name doesn't end with .wav", info.File)
			}
			fullPath := filepath.Join(dir, info.File)
			if !fileutil.FileExists(fullPath) {
				return nil, fmt.Errorf("sample %q can't be located: file does not exist", info.File)
			}
			ctx.Loader.AudioRegistry.Set(id, resource.AudioInfo{
				Path: "$soundbank/" + fullPath,
			})
			fmt.Println("load", "$soundbank/"+fullPath, kind)
			ctx.Loader.LoadAudio(id)

			s := Sample{
				Name:        strings.TrimSuffix(info.File, ".wav"),
				Tag:         instrument.Tag,
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
