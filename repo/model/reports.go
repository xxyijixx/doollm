package model

import (
	"time"
)

const TableNameReport = "pre_reports"

type Report struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	Title     string    `gorm:"column:title;not null;comment:标题" json:"title"`
	Type      string    `gorm:"column:type;not null;default:daily;comment:汇报类型" json:"type"`
	Userid    int64     `gorm:"column:userid;not null" json:"userid"`
	Content   string    `gorm:"column:content" json:"content"`
	Sign      string    `gorm:"column:sign;not null;comment:汇报唯一标识" json:"sign"`
}

func (*Report) TableName() string {
	return TableNameReport
}
