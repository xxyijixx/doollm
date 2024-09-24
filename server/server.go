package server

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	g := gin.Default()
	RouteMount(g)
	StartScheduledTask()
	g.Run(":9090")
}
