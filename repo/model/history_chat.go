package model

const TableNameHistoryChat = "pre_history_aichats"

type HistoryChat struct {
	ID           int64  `json:"id"`
	SessionID    string `gorm:"size:100" json:"session_id"`
	Model        string `gorm:"size:100" json:"model"`
	UserID       string `gorm:"size:100" json:"user_id"`
	LastMessages string `gorm:"size:5000;default:''" json:"last_messages"`
	CreateTime   string `gorm:"type:time" json:"create_time"`
	UpdateTime   string `gorm:"type:time" json:"update_time"`
	Avatar       string `gorm:"size:100" json:"avatar"`
}

func (*HistoryChat) TableName() string {
	return TableNameHistoryChat
}
