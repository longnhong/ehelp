package main

import (
	// 1. init first
	_ "ehelp/init"
	// 2. iniit 2nd
	"ehelp/api"
	"ehelp/middleware"
	"ehelp/room"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(middleware.AddHeader(), gin.Logger(), middleware.Recovery())
	//static
	router.StaticFS("/static", http.Dir("./static"))
	//api
	rootAPI := router.Group("/api")
	api.InitApi(rootAPI)
	//ws
	room.NewRoomServer(router.Group("/room"))
	router.Run(":8080")
}
