package user

import (
	"ehelp/o/user"
	"ehelp/x/rest"

	"github.com/gin-gonic/gin"
)

type UserServer struct {
	*gin.RouterGroup
	rest.JsonRender
}

func NewUserServer(parent *gin.RouterGroup, name string) {
	var s = UserServer{
		RouterGroup: parent.Group(name),
	}
	s.POST("/create", s.handleCreate)
}
func (s *UserServer) handleCreate(ctx *gin.Context) {
	var u *user.Staff
	ctx.BindJSON(&u)
	u.Create()
	s.SendData(ctx, u)
}
