package file

import (
	"testing"
)

func TestQueryFile(t *testing.T) {
	fileService := &FileServiceImpl{}
	fileService.Traversal()
}
