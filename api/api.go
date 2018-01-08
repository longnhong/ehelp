package api

import (
	"ehelp/api/auth"
	"ehelp/api/service"
	"ehelp/api/user"

	"github.com/gin-gonic/gin"
)

func InitApi(root *gin.RouterGroup) {
	service.NewServiceServer(root, "service")
	user.NewUserServer(root, "user")
	auth.NewAuthServer(root, "auth")
}
