package server

import (
	"context"
	"doollm/repo"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Run() {
	r := gin.Default()
	files, err := repo.File.WithContext(context.Background()).Find()
	if err != nil {

	}
	for _, file := range files {
		log.Info("文件信息:", file.ID, file.Name)
	}

	r.Run(":9090")
}
