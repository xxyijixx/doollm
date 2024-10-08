package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameProjectColumn = "pre_project_columns"

// ProjectColumn mapped from table <pre_project_columns>
type ProjectColumn struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProjectID int64          `gorm:"column:project_id;comment:项目ID" json:"project_id"` // 项目ID
	Name      string         `gorm:"column:name;comment:列表名称" json:"name"`             // 列表名称
	Color     string         `gorm:"column:color;comment:颜色" json:"color"`             // 颜色
	Sort      int32          `gorm:"column:sort;comment:排序(ASC)" json:"sort"`          // 排序(ASC)
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName ProjectColumn's table name
func (*ProjectColumn) TableName() string {
	return TableNameProjectColumn
}
