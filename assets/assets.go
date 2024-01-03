package assets

import (
	"embed"
	"io"
	"os"
	"strings"

	"github.com/quasilyte/ge"
)

//go:embed all:_data
var gameAssets embed.FS

func MakeOpenAssetFunc(ctx *ge.Context) func(path string) io.ReadCloser {
	return func(path string) io.ReadCloser {
		if strings.HasPrefix(path, "$soundbank/") {
			path = strings.TrimPrefix(path, "$soundbank/")
			f, err := os.Open(path)
			if err != nil {
				ctx.OnCriticalError(err)
			}
			return f
		}

		f, err := gameAssets.Open("_data/" + path)
		if err != nil {
			ctx.OnCriticalError(err)
		}
		return f
	}
}

func RegisterResources(ctx *ge.Context) {
	registerImageResources(ctx)
	registerAudioResources(ctx)
}
