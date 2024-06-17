package templates

import "embed"

// FS is the embed.FS that contains the templates.

//go:embed *.gohtml
var FS embed.FS
