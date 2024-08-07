package socket

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"server-lu/models"
	"server-lu/services/test"
	"sync"
	"time"
)

var (
	connects = make(map[string]*models.User)
	upgrade  = websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 5 * time.Second,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	lock sync.RWMutex
)

// Upgrade 升级协议
func Upgrade(c *gin.Context) *models.User {
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return nil
	}

	// 获取用户ID
	userId := c.Query("id")

	// 添加 session
	return addSession(userId, conn)
}

// Read 读取ws消息
func Read(user *models.User) {
	defer func() {
		user.Conn.Close()
		removeSession(user.ID)
	}()

	for {
		_, message, err := user.Conn.ReadMessage()
		if err != nil {
			break
		}

		sendData(user, message)
	}
}

// Write 发送客户端消息
func Write(user *models.User) {
	defer user.Conn.Close()
	for {
		select {
		case message := <-user.Send:
			if err := user.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		}
	}
}

// addSession 添加会话
func addSession(id string, conn *websocket.Conn) *models.User {
	lock.Lock()
	defer lock.Unlock()

	user := &models.User{
		ID:       id,
		Conn:     conn,
		Send:     make(chan []byte, 1000),
		Backpack: []string{"Sword", "Shield", "Potion"},
	}
	connects[id] = user
	return connects[id]
}

// getSession 获取会话
func getSession(id string) *models.User {
	lock.RLock()
	defer lock.RUnlock()

	return connects[id]
}

// removeSession 删除会话
func removeSession(id string) {
	lock.Lock()
	defer lock.Unlock()

	delete(connects, id)
}

// sendData 发送消息
func sendData(user *models.User, message []byte) {
	msg := &models.Message{}
	if err := json.Unmarshal(message, msg); err != nil {
		return
	}

	switch msg.Action {
	case "health":
		test.SendHealth(user)
	default:
		log.Printf("Unknown action: %s", msg.Action)
	}
}
