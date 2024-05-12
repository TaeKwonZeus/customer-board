//go:build prod

package router

import (
	"customer-board/web"
	"github.com/charmbracelet/log"
	"io/fs"
)

func getFrontendAssets() fs.FS {
	f, err := fs.Sub(web.EmbedFrontend, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return f
}
