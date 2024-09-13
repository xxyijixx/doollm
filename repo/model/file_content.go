package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameFileContent = "pre_file_contents"

// FileContent mapped from table <pre_file_contents>
type FileContent struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Fid       int64          `gorm:"column:fid;comment:文件ID" json:"fid"`
	Content   string         `gorm:"column:content;comment:内容" json:"content"`
	Text      string         `gorm:"column:text;comment:内容（主要用于文档类型搜索）" json:"text"`
	Size      int64          `gorm:"column:size;comment:大小(B)" json:"size"`
	Userid    int64          `gorm:"column:userid;comment:会员ID" json:"userid"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName FileContent's table name
func (*FileContent) TableName() string {
	return TableNameFileContent
}
