package router

import (
	"github.com/gin-gonic/gin"
	"server-lu/router/socket"
)

func InitRouter() {
	g := gin.Default()

	// http router 启动
	socket.HttpRouter(g)

	// ws router 启动
	socket.WsRouter(g)

	if err := g.Run(":8081"); err != nil {
		panic(err)
	}
}
