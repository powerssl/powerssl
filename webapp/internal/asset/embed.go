package asset

import "embed"

//go:embed css
var CSS embed.FS

//go:embed favicon.ico
var Favicon []byte

//go:embed index.html
var Template embed.FS

//go:embed js
var JS embed.FS
