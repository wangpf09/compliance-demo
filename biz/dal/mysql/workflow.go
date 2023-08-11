package mysql

import (
	"compliance/workflow/biz/hz_gen/workflow_gorm"
	"compliance/workflow/biz/model"
)

// UpdateWorkflowPriority 更新工作流的优先级
func UpdateWorkflowPriority(params workflow_gorm.UpdateWorkflowPriorityRequest) (int64, error) {
	db := DB.Model(&model.WorkflowCustom{})
	tx := db.Where("status = ?", "WAITING")
	if len(params.WorkflowIds) > 0 {
		tx.Where("workflow_id in (?)", params.WorkflowIds)
	}
	if params.StartTime != "" && params.EndTime != "" {
		// 更新的时间范围
		tx = tx.Where("create_at > ?", params.StartTime)
		tx = tx.Where("create_at < ?", params.EndTime)
	}
	if params.WorkflowTplId != "" {
		tx = tx.Where("workflow_tpl_id = ?", params.WorkflowTplId)
	}

	result := tx.Update("priority", params.Priority)
	return result.RowsAffected, result.Error
}
