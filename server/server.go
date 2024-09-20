package server

import (
	"doollm/service/document"
	"doollm/service/file"
	"doollm/service/report"
	"doollm/service/task"
	"doollm/service/workspace"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	documentService := &document.DocumentServiceImpl{}
	_ = documentService
	documentService.RemoveAll()

	workspaceService := &workspace.WorkspaceServiceImpl{}
	_ = workspaceService
	// workspaceService.Create(1)
	// workspaceService.Create(6)
	fileService := &file.FileServiceImpl{}
	_ = fileService
	// fileService.Traversal()
	// fileService.UploadWorkspace()

	reportService := &report.ReportServiceImpl{}
	_ = reportService
	reportService.Traversal()
	reportService.UploadWorkspace()

	taskService := &task.TaskServiceImpl{}
	_ = taskService
	// taskService.Traversal()

	r.Run(":9090")
}
