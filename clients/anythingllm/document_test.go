package anythingllm

import (
	"doollm/clients/anythingllm/documents"
	"encoding/json"
	"fmt"
	"testing"
)

func TestRawText(t *testing.T) {
	client := NewClient()
	params := documents.RawTextParams{
		TextContent: "# Hello World",
		Metadata: documents.RawTextMetadata{
			Title: "Hello.md",
		},
	}
	client.DocumentUploadRowText(params)
}

func TestDocumentUploadFormString(t *testing.T) {
	client := NewClient()
	resp, err := client.DocumentUploadFormString("## Hello World", "helloworld.md", "md")
	if err != nil {
		t.Errorf("Error create document %v", err)
	}
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(jsonData))
}
