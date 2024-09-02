package assets

import (
	"embed"
	"net/http"
)

//go:embed img/* js/*.min* css/*.min* css/et-book/*
var assetsFS embed.FS

func FileSystem() http.FileSystem {
	return http.FS(assetsFS)
}

func FileServer() http.Handler {
	return http.FileServer(FileSystem())
}
