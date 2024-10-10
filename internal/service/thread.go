package service

import (
	"database/sql"
	"doollm/internal/model"
	"doollm/internal/repository"
	dbModel "doollm/repo/model"
	"fmt"
	"log"
	"time"
)

// 获取某个对话窗口的所有历史记录
func FetchChatHistory(workspaceSlug, threadSlug string) ([]model.ChatMessage, error) {

	resp, err := anythingllmClient.GetChatsForWorkspaceThread(workspaceSlug, threadSlug)
	if err != nil {
		return nil, err
	}
	content := make([]model.ChatMessage, len(resp.History))
	for i, h := range resp.History {
		content[i] = model.ChatMessage{
			Content: h.Content,
		}
	}

	return content, nil
}

// 更新最后一条消息并返回所有字段
func UpdateLastMessage(sessionID, lastMessage string) (dbModel.HistoryChat, error) {
	currentTime := time.Now().Format(time.DateTime)
	query := `UPDATE pre_history_aichats SET last_messages = ?, update_time = ? WHERE session_id = ?`
	_, err := repository.DB.Exec(query, lastMessage, currentTime, sessionID)
	if err != nil {
		log.Printf("Error updating last message for session_id %s: %v", sessionID, err)
		return dbModel.HistoryChat{}, fmt.Errorf("error updating last message: %v", err)
	}

	return GetHistoryChat(sessionID)
}

// 根据 sessionID 获取记录
func GetHistoryChat(sessionID string) (dbModel.HistoryChat, error) {
	var chat dbModel.HistoryChat
	query := `SELECT id, session_id, model, user_id, last_messages, create_time, update_time, avatar FROM pre_history_aichats WHERE session_id = ?`
	err := repository.DB.QueryRow(query, sessionID).Scan(&chat.ID, &chat.SessionID, &chat.Model, &chat.UserID, &chat.LastMessages, &chat.CreateTime, &chat.UpdateTime, &chat.Avatar)
	if err != nil {
		if err == sql.ErrNoRows {
			return dbModel.HistoryChat{}, fmt.Errorf("no record found: %v", err)
		}
		return dbModel.HistoryChat{}, fmt.Errorf("error fetching record: %v", err)
	}
	return chat, nil
}

// 根据用户 ID 获取所有对话窗口的最后一条记录, 并按照更新时间降序排列
func FetchChatsByUserID(userID int) ([]dbModel.HistoryChat, error) {
	var chats []dbModel.HistoryChat
	query := `
			SELECT id, session_id, model, user_id, last_messages, create_time, update_time, avatar
			FROM pre_history_aichats WHERE user_id = ?
			ORDER BY update_time DESC
			`
	rows, err := repository.DB.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying chats: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var chat dbModel.HistoryChat
		if err := rows.Scan(&chat.ID, &chat.SessionID, &chat.Model, &chat.UserID, &chat.LastMessages, &chat.CreateTime, &chat.UpdateTime, &chat.Avatar); err != nil {
			return nil, fmt.Errorf("error scanning chat: %v", err)
		}
		chats = append(chats, chat)
	}

	return chats, nil
}

// 根据用户 ID 查询最新 update_time 的 session_id 并返回所有字段
func FetchLatestSessionIDByUserID(userID int) (dbModel.HistoryChat, error) {
	var session dbModel.HistoryChat
	query := `SELECT * FROM pre_history_aichats WHERE user_id = ? ORDER BY update_time DESC LIMIT 1`
	err := repository.DB.QueryRow(query, userID).Scan(&session.ID, &session.SessionID, &session.Model, &session.UserID, &session.LastMessages, &session.CreateTime, &session.UpdateTime, &session.Avatar)
	if err != nil {
		if err == sql.ErrNoRows {
			return dbModel.HistoryChat{}, fmt.Errorf("no session found for user_id: %d", userID)
		}
		return dbModel.HistoryChat{}, fmt.Errorf("error querying latest session ID: %v", err)
	}
	return session, nil
}

// 删除某个会话
func DeleteExternalSession(workspaceSlug, threadSlug string) error {
	err := anythingllmClient.DeleteWorkspaceThread(workspaceSlug, threadSlug)
	if err != nil {
		return err
	}
	return nil
}

// 从数据库中删除该会话
func DeleteSessionFromDatabase(sessionID string) error {
	query := "DELETE FROM pre_history_aichats WHERE session_id = ?"
	_, err := repository.DB.Exec(query, sessionID)
	if err != nil {
		return fmt.Errorf("error deleting session from database: %v", err)
	}
	return nil
}
