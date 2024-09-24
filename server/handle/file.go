package handle

import (
	"doollm/service/file"
	"strconv"

	"github.com/gin-gonic/gin"
)

var fileService = &file.FileServiceImpl{}

func FileShareUpdateHandle(c *gin.Context) {
	var requestMap = make(map[string]string)
	c.Bind(&requestMap)
	if fileId, exists := requestMap["fileId"]; exists {
		fid, err := strconv.ParseInt(fileId, 10, 64)
		if err != nil {
			return
		}
		fileService.Update(fid)
	}
}
