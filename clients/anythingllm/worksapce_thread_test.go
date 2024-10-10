package anythingllm

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateWorkspaceThread(t *testing.T) {
	client := NewClient()
	resp, err := client.CreateWorkspaceThread("workspace-for-user-1")
	if err != nil {
		t.Errorf("Error request error: %v", err)
		return
	}
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(jsonData))
}

func TestGetChatsForWorkspaceThread(t *testing.T) {
	client := NewClient()
	resp, err := client.GetChatsForWorkspaceThread("workspace-for-user-1", "9307a647-1219-4f20-97b0-81fe5e1e0dfc")
	if err != nil {
		t.Errorf("Error get chats for a workspace thread: %v", err)
		return
	}
	jsonData, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(jsonData))
}
