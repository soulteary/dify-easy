package server

import (
	"embed"
	"fmt"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	static "github.com/soulteary/gin-static"
)

func InternalServer(host string, port int, dirRoot string, embedFS embed.FS, embedMode bool) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	if embedMode {
		r.NoRoute(static.ServeEmbed(dirRoot, embedFS))
	} else {
		os.MkdirAll(dirRoot, os.ModePerm)
		r.Use(static.Serve("/", static.LocalFile(dirRoot, true)))
	}

	r.Run(fmt.Sprintf("%s:%d", host, port))
}
