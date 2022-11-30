package server

import (
	"io"
	"log"
	"os"
	"stir/server/handler"

	"github.com/gin-gonic/gin"
)

var Version = "0.1.0"

type ServerOption struct {
	Host string
	Port string
}

func ConfigRouter(r *gin.Engine) {
	api := r.Group("/api")
	api.GET("/ping", handler.PingHandler)
	api.GET("/ip", handler.ClientIPHandler)
	api.GET("/notice", handler.NoticeHandler)
	api.GET("/version", handler.VersionHandler)
	api.GET("/server_status", handler.ServerStatusHandler)
	api.GET("/exchange_rate", handler.ExchangeHandler)

	apiUser := api.Group("/user")
	apiUser.GET("/translation", handler.TranslationHandler)

	// apiV1 := r.Group("/api/v1")

}

func Start(host, port, mode string) error {
	gin.SetMode(mode)
	if mode == gin.ReleaseMode {
		gin.DisableConsoleColor()
		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}
	r := gin.Default()
	ConfigRouter(r)
	log.Println("服务器启动……")
	return r.Run(host + ":" + port) // listen and serve on 0.0.0.0:8080
}
