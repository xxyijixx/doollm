package task

import (
	"time"
)

type TaskRowText struct {
	ProjectName     string           `json:"project_name,omitempty"`
	ProjectOwner    string           `json:"project_owner,omitempty"`
	ProjectUser     string           `json:"project_user,omitempty"`
	ColumnName      string           `json:"column_name,omitempty"`
	TaskName        string           `json:"task_name"`
	TaskOwner       string           `json:"task_owner"`
	TaskAssistant   string           `json:"task_assistant"`
	TaskDescription string           `json:"task_description"`
	Priority        string           `json:"priority,omitempty"`
	Attachment      []string         `json:"attachment,omitempty"`
	Status          string           `json:"status"`
	SubNum          int              `json:"sub_num"`
	SubComplete     int              `json:"sub_complete"`
	SubTask         []SubTaskRowText `json:"sub_task"`
	StartAt         time.Time        `json:"start_at"`
	EndAt           time.Time        `json:"end_at"`
	CompleteAt      time.Time        `json:"complete_at,omitempty"`
}

type SubTaskRowText struct {
	Name       string    `json:"name"`
	Owner      string    `json:"owner"`
	Status     string    `json:"status"`
	StartAt    time.Time `json:"start_at"`
	EndAt      time.Time `json:"end_at"`
	CompleteAt time.Time `json:"complete_at,omitempty"`
}
