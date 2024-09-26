package server

import (
	"doollm/server/schedule"
	"doollm/service/file"
	"doollm/service/report"
	"doollm/service/task"

	log "github.com/sirupsen/logrus"
)

var (
	fileService   = &file.FileServiceImpl{}
	reportService = &report.ReportServiceImpl{}
	taskService   = &task.TaskServiceImpl{}
)

func StartScheduledTask() {
	scheduledTask := schedule.NewScheduledTask()

	// 每天零点执行 0 0 0 * * ?
	_, err := scheduledTask.AddTask("0 0,30 * * * ? ", fileService.Traversal)
	if err != nil {
		log.Errorf("开启定时任务失败 %v", err)
	}
	_, err = scheduledTask.AddTask("0 5,35 * * * ? ", taskService.Traversal)
	if err != nil {
		log.Errorf("开启定时任务失败 %v", err)
	}
	_, err = scheduledTask.AddTask("0 10,40 * * * ? ", reportService.Traversal)
	if err != nil {
		log.Errorf("开启定时任务失败 %v", err)
	}
	scheduledTask.Start()
}
