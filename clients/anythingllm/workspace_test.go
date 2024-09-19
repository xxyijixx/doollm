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
