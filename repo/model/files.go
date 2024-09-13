package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameFile = "pre_files"

type File struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Pid       int64          `gorm:"column:pid;comment:上级ID" json:"pid"`
	Pids      string         `gorm:"column:pids;comment:上级ID递归" json:"pids"`
	Cid       int64          `gorm:"column:cid;comment:复制ID" json:"cid"`
	Name      string         `gorm:"column:name;comment:名称" json:"name"`
	Type      string         `gorm:"column:type;comment:类型" json:"type"`
	Ext       string         `gorm:"column:ext;comment:后缀名" json:"ext"`
	Size      int64          `gorm:"column:size;comment:大小(B)" json:"size"`
	Userid    int64          `gorm:"column:userid;comment:拥有者ID" json:"userid"`
	Share     int32          `gorm:"column:share;comment:是否共享" json:"share"`
	Pshare    int64          `gorm:"column:pshare;comment:所属分享ID" json:"pshare"`
	CreatedID int64          `gorm:"column:created_id;comment:创建者" json:"created_id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (*File) TableName() string {
	return TableNameFile
}
