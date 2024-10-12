package model

import "time"

const TableNameWorkspacePermission = "pre_workspace_permissions"

type WorkspacePermission struct {
	ID          uint64    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID      uint64    `gorm:"column:user_id;not null;comment:同步pre_users表" json:"user_id"`
	IsCreate    bool      `gorm:"column:is_create;default:false" json:"is_create"`
	WorkspaceID string    `gorm:"column:workspace_id;size:100;comment:传入工作区的slug" json:"workspace_id"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
}

func (*WorkspacePermission) TableName() string {
	return TableNameWorkspacePermission
}
