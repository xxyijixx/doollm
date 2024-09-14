package model

import (
	"time"
)

const TableNameLlmDocument = "pre_llm_document"

type LlmDocument struct {
	ID                 int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	LinkType           string    `gorm:"column:link_type;comment:关联类型" json:"link_type"`
	LinkId             int64     `gorm:"column:link_id;comment:关联ID" json:"link_id"`
	LinkParantId       int64     `gorm:"column:link_parant_id;comment:关联ID" json:"link_parant_id"`
	LinkExtras         string    `gorm:"column:link_extras;comment:关联拓展信息" json:"link_extras"`
	DocID              string    `gorm:"column:doc_id;comment:文档ID" json:"doc_id"`
	Name               string    `gorm:"column:name;comment:名称" json:"name"`
	Title              string    `gorm:"column:title;comment:标题" json:"title"`
	Location           string    `gorm:"column:location;comment:位置" json:"location"`
	TokenCountEstimate int64     `gorm:"column:token_count_estimate;comment:预估Token数" json:"token_count_estimate"`
	Userid             int64     `gorm:"column:userid;comment:拥有者ID" json:"userid"`
	LastModifiedAt     time.Time `gorm:"column:last_modified_at;上次修改时间" json:"last_modified_at"`
	CreatedAt          time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (*LlmDocument) TableName() string {
	return TableNameLlmDocument
}
