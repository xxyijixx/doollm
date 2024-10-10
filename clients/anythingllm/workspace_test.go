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
		Name: "workspace-for-user-5",
	})
	if err != nil {
		fmt.Printf("Error create workspace %v", err)
		return

	}
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(jsonData))
}

func TestDeleteWorkspace(t *testing.T) {
	client := NewClient()
	err := client.DeleteWorkspace("workspace-for-user-4")
	if err != nil {
		t.Errorf("err: %v", err)
	}
}

func TestCreateAndDeleteWorkspace(t *testing.T) {
	client := NewClient()
	slug := ""
	t.Run("CreateWorkspace", func(t *testing.T) {
		resp, err := client.CreateWorkspace(workspace.CreateParams{
			Name: "workspace-for-user-5",
		})
		if err != nil {
			t.Errorf("Error create workspace %v", err)
			return
		}
		slug = resp.Workspace.Slug
	})

	t.Run("DeleteWorkspace", func(t *testing.T) {
		if slug == "" {
			t.Error("Error delete workspace, slug is empty")
		}
		err := client.DeleteWorkspace(slug)
		if err != nil {
			t.Errorf("err: %v", err)
		}
	})
}

func TestQueryWorkspaces(t *testing.T) {
	client := NewClient()
	resp, err := client.QueryWorkspaces()
	if err != nil {
		fmt.Printf("Error create workspace %v", err)
		t.Errorf("Error create workspace %v", err)
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
