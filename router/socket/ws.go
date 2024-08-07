package socket

import (
	"github.com/gin-gonic/gin"
	"server-lu/handler/ws"
)

func WsRouter(g *gin.Engine) {
	g.GET("/ws", ws.SingletonWsHandler().InitHandler)
}
