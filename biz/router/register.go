// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"compliance/workflow/biz/handler/workflow_gorm"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	r.POST("/update/workflow", workflow_gorm.UpdateWorkflowPriority)
}
