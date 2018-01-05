package api

import (
	"ehelp/api/service"

	"github.com/gin-gonic/gin"
)

func InitApi(root *gin.RouterGroup) {
	service.NewServiceServer(root, "service")
}
