package service

import (
	"ehelp/o/service"
	_ "ehelp/o/user"
	"ehelp/x/rest"
	"g/x/web"

	"github.com/gin-gonic/gin"
)

type ServiceServer struct {
	*gin.RouterGroup
	rest.JsonRender
}

func NewServiceServer(parent *gin.RouterGroup, name string) {
	var s = &ServiceServer{
		RouterGroup: parent.Group(name),
	}
	s.POST("/create", s.handleCreate)
}

func (s *ServiceServer) handleCreate(ctx *gin.Context) {
	var srv *service.Service
	web.AssertNil(ctx.ShouldBindJSON(&srv))
	web.AssertNil(service.Create(srv))

	s.SendData(ctx, srv)
}
