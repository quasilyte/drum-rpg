package session

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/quasilyte/drum-rpg/contentdir"
	"github.com/quasilyte/drum-rpg/fileutil"
	"github.com/quasilyte/drum-rpg/jsonc"
)

type ContentDirManager struct {
	Dir string

	drumkitsDir string
}

func (m *ContentDirManager) Init() error {
	m.drumkitsDir = filepath.Join(m.Dir, "drumkits")

	if !fileutil.FileExists(m.drumkitsDir) {
		if err := os.Mkdir(m.drumkitsDir, os.ModePerm); err != nil {
			return fmt.Errorf("create drumkits folder: %w", err)
		}
	}

	return nil
}

func (m *ContentDirManager) ListDrumKits() []string {
	var names []string
	m.WalkDrumKits(func(filename string, kit *contentdir.DrumKit) {
		names = append(names, kit.Name)
	})
	sort.Strings(names)
	return names
}

func (m *ContentDirManager) WalkDrumKits(visit func(filename string, kit *contentdir.DrumKit)) {
	files, err := os.ReadDir(m.drumkitsDir)
	if err != nil {
		return
	}

	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".jsonc") {
			continue
		}
		fullFilename := filepath.Join(m.drumkitsDir, f.Name())
		data, err := os.ReadFile(fullFilename)
		if err != nil {
			continue
		}
		parsed := &contentdir.DrumKit{}
		if err := jsonc.Unmarshal(data, parsed); err != nil {
			continue
		}
		visit(fullFilename, parsed)
	}
}
