package auth

import (
	"ehelp/o/auth"
	"ehelp/o/user"
	"ehelp/x/rest"

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
		Role     string `json:"role"`
	}{}
	ctx.BindJSON(&loginInfo)
	u, err := user.GetByUNamePwd(loginInfo.UName, loginInfo.Password, loginInfo.Role)
	rest.AssertNil(err)
	auth, err := auth.Create(u.ID, string(u.Role))
	rest.AssertNil(err)
	s.SendData(ctx, map[string]interface{}{
		"access_token": auth.ID,
	})
}
