package model

import (
	"time"
)

const TableNameLlmWorkspaceDocument = "pre_llm_workspace_document"

// LlmWorkspaceDocument mapped from table <pre_llm_workspace_document>
type LlmWorkspaceDocument struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	WorkspaceID int64     `gorm:"column:workspace_id" json:"workspace_id"`
	DocumentID  int64     `gorm:"column:document_id;comment:文档ID" json:"document_id"` // 文档ID
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
}

// TableName LlmWorkspaceDocument's table name
func (*LlmWorkspaceDocument) TableName() string {
	return TableNameLlmWorkspaceDocument
}
