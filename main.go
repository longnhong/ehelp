package main

import (
	"ehelp/api"
	_ "ehelp/init"
	"ehelp/middleware"
	"ehelp/room"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger(), middleware.Recovery(), middleware.AddHeader())
	//static
	router.StaticFS("/static", http.Dir("./static"))
	//api
	rootAPI := router.Group("/api")
	api.InitApi(rootAPI)
	//ws
	room.NewRoomServer(router.Group("/room"))
	router.Run(":8080")
}
