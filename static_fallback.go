//go:build !embed
// +build !embed

package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

// when not embedding, serve from disk
var frontendFS fs.FS = os.DirFS(filepath.Join(".", "frontend", "dist"))
