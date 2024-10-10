package workspacethread

import "time"

// 创建返回响应
type CreateResponse struct {
	Thread struct {
		ID            int       `json:"id"`
		Name          string    `json:"name"`
		Slug          string    `json:"slug"`
		WorkspaceID   int       `json:"workspace_id"`
		UserID        int       `json:"user_id"`
		CreatedAt     time.Time `json:"createdAt"`
		LastUpdatedAt time.Time `json:"lastUpdatedAt"`
	} `json:"thread"`
	Message interface{} `json:"message"`
}

type HistoryChatsResponse struct {
	History []struct {
		Role          string        `json:"role"`
		Content       string        `json:"content"`
		SentAt        int           `json:"sentAt"`
		Attachments   []interface{} `json:"attachments,omitempty"`
		ChatID        int           `json:"chatId"`
		Type          string        `json:"type,omitempty"`
		Sources       []interface{} `json:"sources,omitempty"`
		FeedbackScore interface{}   `json:"feedbackScore,omitempty"`
	} `json:"history"`
}
