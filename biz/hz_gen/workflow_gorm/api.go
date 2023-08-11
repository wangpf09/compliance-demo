package workflow_gorm

type UpdateWorkflowPriorityRequest struct {
	StartTime     string   `json:"StartTime"`
	EndTime       string   `json:"EndTime"`
	WorkflowIds   []string `json:"WorkflowIds"`
	WorkflowTplId string   `json:"WorkflowTplId"`
	Priority      int      `json:"Priority"`
	DryRun        bool     `json:"DryRun"`
}

type UpdateWorkflowPriorityResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
