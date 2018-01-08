package auth

import (
	"ehelp/o/auth"
	"ehelp/o/user"
	"ehelp/x/rest"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AuthServer struct {
	*gin.RouterGroup
	rest.JsonRender
}

func NewAuthServer(parent *gin.RouterGroup, name string) {
	var s = AuthServer{
		RouterGroup: parent.Group(name),
	}
	s.POST("/signin", s.handleSignin)
}

func (s *AuthServer) handleSignin(ctx *gin.Context) {
	var loginInfo = struct {
		UName    string `json:"uname"`
		Password string `json:"password"`
	}{}
	ctx.BindJSON(&loginInfo)
	u := user.GetByUNamePwd(loginInfo.UName, loginInfo.Password)
	fmt.Println(u)
	res := auth.Create(u.ID, string(u.Password))
	s.SendData(ctx, res)
}
