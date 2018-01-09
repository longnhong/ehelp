package service

import (
	"ehelp/middleware"
	"ehelp/o/service"
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
	s.Use(middleware.MustBeSuperAdmin())
	s.POST("/create", s.handleCreate)
	s.GET("/list", s.handleList)
	s.POST("/tool/create", s.handleCreateTool)
	s.GET("/tool/list", s.handleListTool)
}

func (s *ServiceServer) handleCreate(ctx *gin.Context) {
	var srv *service.Service
	web.AssertNil(ctx.ShouldBindJSON(&srv))
	web.AssertNil(srv.Create())
	s.SendData(ctx, srv)
}
func (s *ServiceServer) handleList(ctx *gin.Context) {
	services, err := service.GetServices()
	rest.AssertNil(err)
	s.SendData(ctx, services)
}
func (s *ServiceServer) handleCreateTool(ctx *gin.Context) {
	var tool *service.Tool
	web.AssertNil(ctx.ShouldBindJSON(&tool))
	web.AssertNil(tool.Create())
	s.SendData(ctx, tool)
}
func (s *ServiceServer) handleListTool(ctx *gin.Context) {
	tools, err := service.GetTools()
	rest.AssertNil(err)
	s.SendData(ctx, tools)
}
