package ws

import (
	"github.com/gin-gonic/gin"
	"server-lu/services/socket"
	"sync"
)

var (
	_wsOnce    sync.Once
	_wsHandler *WsHandler
)

type WsHandler struct{}

func SingletonWsHandler() *WsHandler {
	_wsOnce.Do(func() {
		_wsHandler = &WsHandler{}
	})

	return _wsHandler
}

func (w *WsHandler) InitHandler(c *gin.Context) {
	// 协议升级
	user := socket.Upgrade(c)

	// 开启读取 Goroutine
	go socket.Read(user)

	// 开启发送 Goroutine
	go socket.Write(user)
}
