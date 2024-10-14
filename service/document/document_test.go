package document

import "testing"

var documentService = &DocumentServiceImpl{}

func TestClearDocument(t *testing.T) {
	err := documentService.Clear()
	if err != nil {
		t.Errorf("Error %v", err)
	}
}
