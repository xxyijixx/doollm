// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameProjectUser = "pre_project_users"

// ProjectUser mapped from table <pre_project_users>
type ProjectUser struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProjectID int64     `gorm:"column:project_id;comment:项目ID" json:"project_id"` // 项目ID
	Userid    int64     `gorm:"column:userid;comment:成员ID" json:"userid"`         // 成员ID
	Owner     int32     `gorm:"column:owner;comment:是否负责人" json:"owner"`          // 是否负责人
	TopAt     time.Time `gorm:"column:top_at;comment:置顶时间" json:"top_at"`         // 置顶时间
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName ProjectUser's table name
func (*ProjectUser) TableName() string {
	return TableNameProjectUser
}