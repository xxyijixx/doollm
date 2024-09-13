package model

import (
	"time"
)

const TableNameFileUser = "pre_file_users"

type FileUser struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	FileID     int64     `gorm:"column:file_id;comment:项目ID" json:"file_id"`
	Userid     int64     `gorm:"column:userid;comment:成员ID" json:"userid"`
	Permission int32     `gorm:"column:permission;comment:权限：0只读，1读写" json:"permission"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (*FileUser) TableName() string {
	return TableNameFileUser
}
