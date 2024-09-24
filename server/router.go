package server

import (
	"doollm/server/handle"

	"github.com/gin-gonic/gin"
)

func RouteMount(g *gin.Engine) {
	g.POST("/file/update", handle.FileShareUpdateHandle)
}
