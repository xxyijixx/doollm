package model

import "time"

const TableNameWorkspaceSubscription = "pre_workspace_subscriptions"

type WorkspaceSubscription struct {
	Type      string    `gorm:"column:type;type:enum('free','pro')" json:"type"`
	StartTime time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime   time.Time `gorm:"column:end_time" json:"end_time"`
	IsForever bool      `gorm:"column:is_forever;default:false" json:"is_forever"`
}

func (*WorkspaceSubscription) TableName() string {
	return TableNameWorkspaceSubscription
}
