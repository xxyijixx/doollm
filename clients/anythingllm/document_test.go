package anythingllm

import (
	"doollm/clients/anythingllm/documents"
	"encoding/json"
	"fmt"
	"testing"
)

func TestQueryDocument(t *testing.T) {
	client := NewClient()
	resp, err := client.QueryDocument()
	if err != nil {
		t.Errorf("Error query document %v", err)
	}
	for _, folder := range resp.LocalFiles.Items {
		fmt.Printf("当前文件夹: %s\n", folder.Name)
		for _, document := range folder.Items {
			fmt.Printf("当前文档ID: %s,名称: %s\n", document.ID, document.Name)
		}
	}
}

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
