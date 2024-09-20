package server

import (
	"doollm/service/document"
	"doollm/service/workspace"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	documentService := &document.DocumentServiceImpl{}
	_ = documentService
	// documentService.RemoveAll()

	workspaceService := &workspace.WorkspaceServiceImpl{}
	_ = workspaceService
	// workspaceService.Create(1)
	// workspaceService.Create(6)
	StartScheduledTask()
	r.Run(":9090")
}
