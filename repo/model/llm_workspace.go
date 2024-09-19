package model

import (
	"time"
)

const TableNameLlmWorkspace = "pre_llm_workspace"

// LlmWorkspace mapped from table <pre_llm_workspace>
type LlmWorkspace struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Userid    int64     `gorm:"column:userid;comment:用户ID" json:"userid"` // 用户ID
	Name      string    `gorm:"column:name" json:"name"`
	Slug      string    `gorm:"column:slug" json:"slug"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

// TableName LlmWorkspace's table name
func (*LlmWorkspace) TableName() string {
	return TableNameLlmWorkspace
}
