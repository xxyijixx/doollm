package anythingllm

import (
	"doollm/clients/anythingllm/documents"
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {

	client := NewClient()
	data, _ := client.ValidToken()

	fmt.Println("数据：", data)
}

func TestRawText(t *testing.T) {
	client := NewClient()
	params := documents.RawTextParams{
		TextContent: "# Hello World",
		Metadata: documents.RawTextMetadata{
			Title: "Hello.md",
		},
	}
	client.UploadRowText(params)
}
