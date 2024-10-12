package model

import "time"

const TableNameHistoryChat = "pre_history_aichats"

type HistoryChat struct {
	ID           uint64    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID       string    `gorm:"column:user_id;size:100" json:"user_id"`
	SessionID    string    `gorm:"column:session_id;size:100;unique;not null" json:"session_id"`
	Model        string    `gorm:"column:model;size:100" json:"model"`
	LastMessages string    `gorm:"column:last_messages;size:5000;default:''" json:"last_messages"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	Avatar       string    `gorm:"column:avatar;size:100;comment:模型头像" json:"avatar"`
}

func (*HistoryChat) TableName() string {
	return TableNameHistoryChat
}
