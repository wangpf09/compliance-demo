package workflow_gorm

import (
	"compliance/workflow/biz/dal/mysql"
	"compliance/workflow/biz/hz_gen/workflow_gorm"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	low    = 0
	normal = 25
	high   = 50
)

const layout = "2006-01-02 15:04:05"

func UpdateWorkflowPriority(ctx context.Context, c *app.RequestContext) {
	var err error
	var params workflow_gorm.UpdateWorkflowPriorityRequest
	err = c.BindAndValidate(&params)
	if err != nil {
		logrus.Errorf("参数解析异常: %v", err)
		c.JSON(http.StatusOK, &workflow_gorm.UpdateWorkflowPriorityResponse{
			Code:    1,
			Message: err.Error(),
		})
		return
	}

	if len(params.WorkflowIds) == 0 && (params.StartTime == "" && params.EndTime == "") {
		c.JSON(http.StatusOK, &workflow_gorm.UpdateWorkflowPriorityResponse{
			Code:    1,
			Message: "待更新的workflow为空，并且不存在更新的时间范围",
		})
		return
	}
	satisfy, err := checkTimeRange(params.StartTime, params.EndTime)
	if err != nil {
		logrus.Errorf("解析时间异常: %v", err)
		c.JSON(http.StatusOK, &workflow_gorm.UpdateWorkflowPriorityResponse{
			Code:    1,
			Message: err.Error(),
		})
		return
	}
	if !satisfy {
		c.JSON(http.StatusOK, &workflow_gorm.UpdateWorkflowPriorityResponse{
			Code:    1,
			Message: "时间范围需要在30min内，并且结束时间不能大于开始时间",
		})
		return
	}
	if !(params.Priority == low || params.Priority == normal || params.Priority == high) {
		c.JSON(http.StatusOK, &workflow_gorm.UpdateWorkflowPriorityResponse{
			Code:    1,
			Message: "优先级只允许为: 0、25、50",
		})
		return
	}

	rows, err := mysql.UpdateWorkflowPriority(params)
	if err != nil {
		logrus.Errorf("数据更新异常: %v", err)
		c.JSON(http.StatusOK, &workflow_gorm.UpdateWorkflowPriorityResponse{
			Code:    1,
			Message: "数据更新异常",
		})
	} else {
		c.JSON(http.StatusOK, &workflow_gorm.UpdateWorkflowPriorityResponse{
			Code:    0,
			Message: fmt.Sprintf("更新%d条数据", rows),
		})
	}
}

// checkTimeRange 最大只允许更新30min以内的数据，并且结束时间大于开始时间
func checkTimeRange(startTime, endTime string) (bool, error) {
	start, err := time.Parse(layout, startTime)
	end, err := time.Parse(layout, endTime)
	if err != nil {
		return false, err
	}
	// 结束时间早于开始时间，或者时间范围超过30分钟
	if end.Before(start) || end.Sub(start).Minutes() > 30 {
		return false, nil
	}
	return true, nil
}
