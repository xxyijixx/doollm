package model

import "time"

const TableNameWorkspacePermission = "pre_workspace_permissions"

type WorkspacePermission struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	IsCreate    bool      `json:"is_create"`
	WorkspaceID string    `json:"workspace_id"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
}

func (*WorkspacePermission) TableName() string {
	return TableNameWorkspacePermission
}
