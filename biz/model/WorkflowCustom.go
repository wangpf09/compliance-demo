package model

import (
	"time"
)

type WorkflowCustom struct {
	Id                uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	TenantId          int64     `gorm:"column:tenant_id;NOT NULL"`
	WorkflowId        string    `gorm:"column:workflow_id;NOT NULL"`
	WorkflowTplId     string    `gorm:"column:workflow_tpl_id;NOT NULL"`
	Type              string    `gorm:"column:type;NOT NULL"`
	Status            string    `gorm:"column:status;NOT NULL"`
	CreateAt          time.Time `gorm:"column:create_at;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdateAt          time.Time `gorm:"column:update_at;default:CURRENT_TIMESTAMP;NOT NULL"`
	LastExecutionTime time.Time `gorm:"column:last_execution_time;default:NULL"`
	ProjectId         string    `gorm:"column:project_id;NOT NULL;comment:'项目ID'"`
	Name              string    `gorm:"column:name;default:"`
	Content           string    `gorm:"column:content;default:NULL"`
	Description       string    `gorm:"column:description;default:"`
	StatusLevel2      string    `gorm:"column:status_level2;default:"`
	BatchWorkflowId   string    `gorm:"column:batch_workflow_id;default:"`
	CreatedBy         string    `gorm:"column:created_by;default:"`
	Priority          int32     `gorm:"column:priority;default:0;NOT NULL"`
}

func (w *WorkflowCustom) TableName() string {
	return "workflow_custom"
}
