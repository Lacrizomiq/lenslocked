package templates

import "embed"

// FS holds the embedded templates for the application to use. It is populated by the go:embed directive.
// The go:embed directive is a new feature in Go 1.16 that allows you to embed files directly into your Go code.
// This is useful for embedding static assets like HTML templates, CSS, and JavaScript files.

//go:embed *.gohtml
var FS embed.FS
