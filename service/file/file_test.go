package file

import (
	"testing"
)

var fileService = &FileServiceImpl{}

func TestQueryFile(t *testing.T) {
	fileService.Traversal()
}

func TestUpdate(t *testing.T) {
	fileService.Update(32)
}
