package server

import (
	"doollm/service/file"
	"doollm/service/report"
	"doollm/service/task"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	fileService := &file.FileServiceImpl{}
	_ = fileService
	reportService := &report.ReportServiceImpl{}
	_ = reportService
	// reportService.Traversal()
	// fileService.Traversal()
	taskService := &task.TaskServiceImpl{}
	taskService.Traversal()
	r.Run(":9090")
}
