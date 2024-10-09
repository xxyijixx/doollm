package model

const TableNameHistoryChat = "pre_history_aichats"

type HistoryChat struct {
	ID           int64  `json:"id"`
	SessionID    string `json:"session_id"`
	Model        string `json:"model"`
	UserID       string `json:"user_id"`
	LastMessages string `json:"last_messages"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
	Avatar       string `json:"avatar"`
}

func (*HistoryChat) TableName() string {
	return TableNameHistoryChat
}
