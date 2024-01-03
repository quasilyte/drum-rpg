package tracker

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/quasilyte/drum-rpg/edrum"
	"github.com/quasilyte/xm/xmfile"
)

type Track struct {
	Name   string
	Album  string
	Author string

	Module *xmfile.Module

	instrumentMap [32]edrum.InstrumentKind
}

func (t *Track) GetInstrumentKind(instrumentID int) edrum.InstrumentKind {
	return t.instrumentMap[instrumentID&0x1f]
}

func ParseTrack(album, filename string) (*Track, error) {
	xmBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	metadataFilename := strings.TrimSuffix(filename, filepath.Ext(filename)) + ".json"
	metadataBytes, err := os.ReadFile(metadataFilename)
	if err != nil {
		return nil, err
	}
	var metadata struct {
		Name        string
		Author      string
		Instruments map[string]int
	}
	if err := json.Unmarshal(metadataBytes, &metadata); err != nil {
		return nil, err
	}

	xmParser := xmfile.NewParser(xmfile.ParserConfig{})
	m, err := xmParser.ParseFromBytes(xmBytes)
	if err != nil {
		return nil, err
	}

	t := &Track{
		Name:   metadata.Name,
		Album:  album,
		Author: metadata.Author,
		Module: m,
	}

	for instrumentTag, instrumentID := range metadata.Instruments {
		kind := edrum.InstrumentKindByName(instrumentTag)
		if kind == edrum.UndefinedInstrument {
			return nil, fmt.Errorf("unrecognized instrument type: %q", instrumentTag)
		}
		t.instrumentMap[instrumentID] = kind
	}

	return t, nil
}
