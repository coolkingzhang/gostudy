package main

import (
	"mishop/app/core"
	"mishop/app/routes"

	"github.com/DeanThompson/ginpprof"
	logger "github.com/alecthomas/log4go"
	"github.com/gin-gonic/gin"
)

var (
	addr      = core.Config.String("server", "addr", "")
	logconfig = core.Config.String("log", "logconfig", "")
)

func main() {
	core.Config.Parse()
	logger.LoadConfiguration(*logconfig)
	r := gin.New()
	//	r.Use(core.SessionGin())
	r.LoadHTMLGlob("app/views/**/*")
	r.Static("/assets", "./public/assets")
	r.StaticFile("/favicon.ico", "./public/assets/image/favicon.ico")
	routes.InitRourte(r)

	ginpprof.Wrapper(r)
	r.Run(*addr)
}
