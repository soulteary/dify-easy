package server

import (
	"embed"

	"github.com/gin-gonic/gin"
)

func Launch(embedFS embed.FS) {
	gin.SetMode(gin.ReleaseMode)
	InternalServer("0.0.0.0", 6789, "web", embedFS, true)
}
