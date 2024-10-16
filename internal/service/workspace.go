package service

import (
	"database/sql"
	"doollm/clients/anythingllm/workspace"
	"doollm/internal/repository"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// 发送请求到外部 API 创建工作区
func CreateExternalWorkspace(name string) (string, error) {
	resp, err := anythingllmClient.CreateWorkspace(workspace.CreateParams{
		Name: name,
	})
	if err != nil {
		return "", err
	}
	return resp.Workspace.Slug, nil
}

// 更新 workspace_id
func UpdateWorkspaceID(userID int, slug string) error {
	query := `UPDATE pre_workspace_permissions SET workspace_id = ? WHERE user_id = ?`
	_, err := repository.DB.Exec(query, slug, userID)
	if err != nil {
		return fmt.Errorf("failed to update workspace_id: %v", err)
	}
	return nil
}

// 删除工作区
func DeleteExternalWorkspace(slug string) error {
	err := anythingllmClient.DeleteWorkspace(slug)

	return err
}

// 重置 workspace_id 为 NULL
func ResetWorkspaceID(userID int) error {
	query := `UPDATE pre_workspace_permissions SET workspace_id = NULL WHERE user_id = ?`
	_, err := repository.DB.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("failed to reset workspace_id: %v", err)
	}
	return nil
}

// 根据 userID 获取 workspace_id (slug)
func GetWorkspaceSlug(userID int) (string, error) {
	var slug string
	query := `SELECT workspace_id FROM pre_workspace_permissions WHERE user_id = ?`
	err := repository.DB.QueryRow(query, userID).Scan(&slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no workspace found for user ID %d", userID)
		}
		return "", fmt.Errorf("error querying workspace slug: %v", err)
	}
	return slug, nil
}

// 存储对话窗口 slug 、模型和头像
func StoreChatData(model, avatar, threadSlug string, userID int) error {
	currentTime := time.Now().Format(time.DateTime)
	query := `
	INSERT INTO pre_history_aichats (model, avatar, session_id, user_id, create_time, update_time)
	VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := repository.DB.Exec(query, model, avatar, threadSlug, userID, currentTime, currentTime)
	if err != nil {
		log.Printf("Error storing chat data: %v", err)
		return err
	}
	return nil
}

// 创建新对话窗口
func CreateNewThread(slug string) (string, error) {
	resp, err := anythingllmClient.CreateWorkspaceThread(slug)
	if err != nil {
		return "", nil
	}
	return resp.Thread.Slug, nil
}

// 从对话窗口的 slug 中提取用户 ID
func ExtractUserID(slug string) (int, error) {
	parts := strings.Split(slug, "-")
	if len(parts) < 4 {
		return 0, fmt.Errorf("invalid slug format")
	}
	userID, err := strconv.Atoi(parts[3])
	if err != nil {
		return 0, fmt.Errorf("failed to convert user ID: %v", err)
	}
	return userID, nil
}

func SelectModel(model string) (string, error) {
	parts := strings.Split(model, ",")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid model format, expected 'workspaceSlug, provider, model'")
	}
	workspaceSlug := strings.TrimSpace(parts[0])
	chatProvider := strings.TrimSpace(parts[1])
	chatModel := strings.TrimSpace(parts[2])

	resp, err := anythingllmClient.SelectWorkspaceModel(workspaceSlug, chatProvider, chatModel)
	if err != nil {
		return "select model failed", err
	}

	return resp.Content, nil
}
