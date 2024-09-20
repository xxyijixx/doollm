package anythingllm

import (
	"doollm/clients/anythingllm/workspace"
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateWorkspace(t *testing.T) {
	client := NewClient()
	resp, err := client.CreateWorkspace(workspace.CreateParams{
		Name: "jintianxingqisi",
	})
	if err != nil {
		fmt.Printf("Error create workspace %v", err)
		return

	}

	fmt.Printf("response : %+v", resp)
}

func TestQueryWorkspaces(t *testing.T) {
	client := NewClient()
	resp, err := client.QueryWorkspaces()
	if err != nil {
		fmt.Printf("Error create workspace %v", err)
		return
	}

	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(jsonData))
}

func TestUpdateEmbeddings(t *testing.T) {
	client := NewClient()
	slug := "tttttttttttttttt"
	resp, err := client.UpdateEmbeddings(slug, workspace.UpdateEmbeddingsParams{
		Adds: []string{"custom-documents/raw-report-6-3-1726740643-83d5496b-ac47-4146-ab7a-d3c5d0992aac.json"},
	})
	if err != nil {
		fmt.Println("error ", err)
		return
	}
	fmt.Printf("resp: %v", resp)

}
