//go:build embed
// +build embed

package main

import "embed"

//go:embed frontend/dist/*
var frontendFS embed.FS
