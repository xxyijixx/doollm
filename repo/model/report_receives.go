package model

import (
	"time"
)

const TableNameReportReceive = "pre_report_receives"

type ReportReceive struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Rid         int64     `gorm:"column:rid;not null" json:"rid"`
	ReceiveTime time.Time `gorm:"column:receive_time;comment:接收时间" json:"receive_time"`
	Userid      int64     `gorm:"column:userid;not null;comment:接收人" json:"userid"`
	Read        int32     `gorm:"column:read;not null;comment:是否已读" json:"read"`
}

func (*ReportReceive) TableName() string {
	return TableNameReportReceive
}
