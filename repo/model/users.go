package model

import (
	"time"
)

const TableNameUser = "pre_users"

type User struct {
	Userid       int64     `gorm:"column:userid;primaryKey;autoIncrement:true" json:"userid"`
	Identity     string    `gorm:"column:identity" json:"identity"`
	Department   string    `gorm:"column:department" json:"department"`
	Az           string    `gorm:"column:az;comment:A-Z" json:"az"`
	Pinyin       string    `gorm:"column:pinyin;comment:拼音（主要用于搜索）" json:"pinyin"`
	Email        string    `gorm:"column:email" json:"email"`
	Tel          string    `gorm:"column:tel;comment:联系电话" json:"tel"`
	Nickname     string    `gorm:"column:nickname" json:"nickname"`
	Profession   string    `gorm:"column:profession" json:"profession"`
	Userimg      string    `gorm:"column:userimg" json:"userimg"`
	Encrypt      string    `gorm:"column:encrypt" json:"encrypt"`
	Password     string    `gorm:"column:password;comment:登录密码" json:"password"`
	Changepass   int32     `gorm:"column:changepass;comment:登录需要修改密码" json:"changepass"`
	LoginNum     int32     `gorm:"column:login_num;comment:累计登录次数" json:"login_num"`
	LastIP       string    `gorm:"column:last_ip;comment:最后登录IP" json:"last_ip"`
	LastAt       time.Time `gorm:"column:last_at;comment:最后登录时间" json:"last_at"`
	LineIP       string    `gorm:"column:line_ip;comment:最后在线IP（接口）" json:"line_ip"`
	LineAt       time.Time `gorm:"column:line_at;comment:最后在线时间（接口）" json:"line_at"`
	TaskDialogID int64     `gorm:"column:task_dialog_id;comment:最后打开的任务会话ID" json:"task_dialog_id"`
	CreatedIP    string    `gorm:"column:created_ip;comment:注册IP" json:"created_ip"`
	DisableAt    string    `gorm:"column:disable_at" json:"disable_at"`
	EmailVerity  bool      `gorm:"column:email_verity;comment:邮箱是否已验证" json:"email_verity"`
	Bot          int32     `gorm:"column:bot;comment:是否机器人" json:"bot"`
	CreatedAt    string    `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (*User) TableName() string {
	return TableNameUser
}
