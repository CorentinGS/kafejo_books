package assets

import (
	"embed"
	"net/http"
)

//go:embed img/* js/*.min* css/*.min* css/et-book/*
var assetsFS embed.FS

//go:embed public/robots.txt
var robotsTxt []byte

//go:embed public/sitemap.xml
var sitemapXML []byte

func FileSystem() http.FileSystem {
	return http.FS(assetsFS)
}

func FileServer() http.Handler {
	return http.FileServer(FileSystem())
}

func RobotsTxt() []byte {
	return robotsTxt
}

func SitemapXML() []byte {
	return sitemapXML
}
