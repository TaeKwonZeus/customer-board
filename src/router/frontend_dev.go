//go:build !prod

package router

import (
	"io/fs"
	"os"
)

func getFrontendAssets() fs.FS {
	return os.DirFS("web/dist")
}
