//go:build prod

package web

import "embed"

//go:embed dist
var EmbedFrontend embed.FS
